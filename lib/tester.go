package lib

import (
	"encoding/csv"
	"fmt"
	"loadtest/models"
	"log"
	"os"
	"strconv"
	"sync"
)

type Response struct {
	responseCode  int
	responseTime  float64
	firstByteTime float64
	lastByteTime  float64
}

type Test struct {
	url         string
	requests    int
	concurrency int
}

func TestUrl(url string, requests int, concurrency int) {

	fmt.Println("Testing url: ", url)
	fmt.Println("Number of requests: ", requests)

	analysis_channel := make(chan *models.Analysis)
	test_channel := make(chan *Test)

	var wg sync.WaitGroup
	wg.Add(1)
	go testUrlWorker(0, test_channel, analysis_channel, &wg)

	test := Test{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
	}
	test_channel <- &test

	analysis := <-analysis_channel

	fmt.Println()
	models.DisplayAnalysis(analysis)
}

func TestFile(filepath string, concurrent_urls int) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records")
	}

	analysis_channel := make(chan *models.Analysis)
	test_channel := make(chan *Test)

	var wg sync.WaitGroup
	for i := 0; i < concurrent_urls; i++ {
		wg.Add(1)
		go testUrlWorker(i, test_channel, analysis_channel, &wg)
	}

	go displayAnalysisWorker(analysis_channel)

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

		test := Test{
			url:         url,
			requests:    requests,
			concurrency: concurrency,
		}
		log.Println("Adding to queue: ", test)
		test_channel <- &test
	}
	close(test_channel)
	wg.Wait()
	analysis_channel <- nil
}
