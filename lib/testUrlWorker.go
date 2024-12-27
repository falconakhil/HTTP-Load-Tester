package lib

import (
	"loadtest/models"
	"log"
	"sync"
	"time"
)

func testUrlWorker(workerID int, test_channel chan *Test, analysis_channel chan *models.Analysis, wg *sync.WaitGroup) {

	defer wg.Done()
	for test := range test_channel {
		log.Printf("testUrlWorker %d processing job: %#v", workerID, test)
		url := test.url
		requests := test.requests
		concurrency := test.concurrency

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
}
