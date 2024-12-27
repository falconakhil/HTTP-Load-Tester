package lib

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logFile, _ = os.OpenFile("/dev/null", os.O_RDWR, 0644)

type Response struct {
	responseCode  int
	responseTime  float64
	firstByteTime float64
	lastByteTime  float64
}

func TestUrl(url string, requests int, concurrency int) {
	startTime := time.Now()
	log.SetOutput(logFile)
	response := make(chan Response, requests)
	jobs := make(chan int)
	log.Println("Requests", requests)
	for i := 0; i < concurrency; i++ {
		go sendRequestWorker(i, url, jobs, response)
	}

	for i := 0; i < requests; i++ {
		jobs <- i
	}
	log.Println("Closing jobs")
	close(jobs)
	analyzeResponses(requests, response, startTime)
}

func analyzeResponses(requests int, response chan Response, startTime time.Time) {
	success := 0
	failure := 0

	firstByteMin := 1000000.0
	firstByteMax := 0.0
	firstByteAvg := 0.0

	lastByteMax := 0.0
	lastByteMin := 1000000.0
	lastByteAvg := 0.0

	responseTimeMax := 0.0
	responseTimeMin := 1000000.0
	responseTimeAvg := 0.0

	for i := 0; i < requests; i++ {
		resp := <-response
		if resp.responseCode == 200 {
			success = success + 1

			// First Byte times
			firstByteMax = max(firstByteMax, resp.firstByteTime)
			firstByteMin = min(firstByteMin, resp.firstByteTime)
			firstByteAvg = firstByteAvg + resp.firstByteTime

			// Last Byte times
			lastByteMax = max(lastByteMax, resp.lastByteTime)
			lastByteMin = min(lastByteMin, resp.lastByteTime)
			lastByteAvg = lastByteAvg + resp.lastByteTime

			// Response times
			responseTimeMax = max(responseTimeMax, resp.responseTime)
			responseTimeMin = min(responseTimeMin, resp.responseTime)
			responseTimeAvg = responseTimeAvg + resp.responseTime
		} else {
			failure = failure + 1
		}

	}
	endTime := time.Since(startTime)
	close(response)
	fmt.Println("Analysis:")
	fmt.Println("------------------------------------------")
	fmt.Println("Successful: ", success)
	fmt.Println("Failed: ", failure)
	fmt.Println("------------------------------------------")
	fmt.Println("Requests per second", float64(requests)/endTime.Seconds())
	fmt.Println("------------------------------------------")
	fmt.Println("First Byte Time", firstByteMin, firstByteMax, firstByteAvg/float64(success))
	fmt.Println("Last Byte Time", lastByteMin, lastByteMax, lastByteAvg/float64(success))
	fmt.Println("Response Time", responseTimeMin, responseTimeMax, responseTimeAvg/float64(success))

}
