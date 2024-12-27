package models

import "fmt"

type Analysis struct {
	url         string
	requests    int
	concurrency int

	successful        int
	failed            int
	requestsPerSecond float64

	firstByteMin float64
	firstByteMax float64
	firstByteAvg float64

	lastByteMin float64
	lastByteMax float64
	lastByteAvg float64

	responseTimeMin float64
	responseTimeMax float64
	responseTimeAvg float64
}

func InitalizeAnalysis(url string, requests int, concurrenct int) Analysis {
	return Analysis{
		url:         url,
		requests:    requests,
		concurrency: concurrenct,
	}
}

func AddData(
	analysis *Analysis,
	successful int,
	failed int,
	requestsPerSecond float64,
	firstByteMin float64,
	firstByteMax float64,
	firstByteAvg float64,
	lastByteMin float64,
	lastByteMax float64,
	lastByteAvg float64,
	responseTimeMin float64,
	responseTimeMax float64,
	responseTimeAvg float64) {

	analysis.successful = successful
	analysis.failed = failed
	analysis.requestsPerSecond = requestsPerSecond

	analysis.firstByteMin = firstByteMin
	analysis.firstByteMax = firstByteMax
	analysis.firstByteAvg = firstByteAvg

	analysis.lastByteMin = lastByteMin
	analysis.lastByteMax = lastByteMax
	analysis.lastByteAvg = lastByteAvg

	analysis.responseTimeMin = responseTimeMin
	analysis.responseTimeMax = responseTimeMax
	analysis.responseTimeAvg = responseTimeAvg
}

func DisplayAnalysis(analysis *Analysis) {
	fmt.Println("URL: ", analysis.url)
	fmt.Println("Requests: ", analysis.requests)
	fmt.Println("Concurrency: ", analysis.concurrency)
	fmt.Println("------------------------------------------")
	fmt.Println("Successful: ", analysis.successful)
	fmt.Println("Failed: ", analysis.failed)
	fmt.Println("------------------------------------------")
	fmt.Println("Requests per second", analysis.requestsPerSecond)
	fmt.Println("------------------------------------------")
	fmt.Println("Time\t\tMin\t\tMax\t\tAvg\t\t(milliseconds)")
	fmt.Printf("First Byte\t%f\t%f\t%f\n", analysis.firstByteMin, analysis.firstByteMax, analysis.firstByteAvg)
	fmt.Printf("Last Byte\t%f\t%f\t%f\n", analysis.lastByteMin, analysis.lastByteMax, analysis.lastByteAvg)
	fmt.Printf("Response\t%f\t%f\t%f\n", analysis.responseTimeMin, analysis.responseTimeMax, analysis.lastByteAvg)
}
