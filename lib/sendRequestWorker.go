package lib

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func sendRequestWorker(workerID int, url string, jobs chan int, response chan Response) {
	for j := range jobs {
		log.Printf("sendRequestWorker %d processing job %d\n", workerID, j)

		resp := Response{
			responseCode:  -1,
			responseTime:  -1,
			firstByteTime: -1,
			lastByteTime:  -1,
		}

		// Set up a tcp connection
		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Println("Error: ", err)
			response <- resp
			continue
		}

		defer conn.Close()

		// Send a HTTP GET request
		conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
		start := time.Now()

		// Read the first byte
		oneByte := make([]byte, 1)
		_, err = conn.Read(oneByte)
		if err != nil {
			log.Println("Error: ", err)
			response <- resp
			continue
		}
		firstByteTime := time.Since(start)

		// Read the entire response
		http_response, err := io.ReadAll(conn)
		if err != nil {
			log.Println("Error: ", err)
			response <- resp
			continue
		}
		lastByteTime := time.Since(start)

		response_code := parseStatusCode(string(http_response))
		if response_code == -1 {
			log.Println("Error: Couldn't parse status code")
			response <- resp
			continue
		} else if response_code != 200 {
			log.Println("Error: Response code is not 200, but ", response_code)
		}

		resp.responseCode = response_code
		resp.firstByteTime = float64(firstByteTime.Milliseconds())
		resp.lastByteTime = float64(lastByteTime.Milliseconds())
		resp.responseTime = float64(lastByteTime.Milliseconds() - firstByteTime.Milliseconds())
		response <- resp
	}
}

func parseStatusCode(response string) int {
	// The response starts with something like "HTTP/1.1 200 OK"
	parts := strings.Split(response, " ")
	if len(parts) > 1 {
		// Convert the second part to an integer (the status code)
		var statusCode int
		_, err := fmt.Sscanf(parts[1], "%d", &statusCode)
		if err == nil {
			return statusCode
		}
	}
	return -1 // Return -1 if we couldn't parse the status code
}
