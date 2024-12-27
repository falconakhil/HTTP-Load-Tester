package lib

import (
	"encoding/csv"
	"fmt"
	"loadtest/models"
	"log"
	"os"
	"strconv"
	"time"
)

type Response struct {
	responseCode  int
	responseTime  float64
	firstByteTime float64
	lastByteTime  float64
}

func TestUrl(url string, requests int, concurrency int) {
	analysis_channel := make(chan *models.Analysis)
	go testUrl(url, requests, concurrency, analysis_channel)
	analysis := <-analysis_channel
	fmt.Println()
	models.DisplayAnalysis(analysis)
}

func TestFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records")
	}

	analysis_channel := make(chan *models.Analysis)
	urls := 0
	for _, record := range records {
		url := record[0]

		requests, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal("Error while converting request to int", err)
		}

		concurrency, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal("Error while converting concurrency to int", err)
		}

		go testUrl(url, requests, concurrency, analysis_channel)
		urls = urls + 1
	}
	defer file.Close()
	for i := 0; i < urls; i++ {
		analysis := <-analysis_channel
		fmt.Println()
		models.DisplayAnalysis(analysis)
	}
	close(analysis_channel)

}

func testUrl(url string, requests int, concurrency int, analysis_channel chan *models.Analysis) {
	fmt.Println("Testing url: ", url)
	fmt.Println("Number of requests: ", requests)

	log.Println("Testing url: ", url)
	log.Println("Number of requests: ", requests)
	log.Println("Concurrency", concurrency)

	response := make(chan Response, requests)
	jobs := make(chan int)
	for i := 0; i < concurrency; i++ {
		go sendRequestWorker(i, url, jobs, response)
	}

	startTime := time.Now()
	for i := 0; i < requests; i++ {
		jobs <- i
	}
	log.Println("Closing jobs")
	close(jobs)

	analysis := models.InitalizeAnalysis(url, requests, concurrency)
	analyzeResponses(requests, response, startTime, &analysis)
	analysis_channel <- &analysis
}
