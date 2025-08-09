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

// AnalysisCollection holds multiple analysis results for different concurrency levels
type AnalysisCollection struct {
    Results []Analysis
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

func AccumulateAsCollection(analysis_channel chan *Analysis,collection *AnalysisCollection)  {
	for analysis := range analysis_channel {
		if(analysis==nil) {
			break
		}
		collection.Results = append(collection.Results, *analysis)

	}
}

// Getter functions for Anaylsis
func (a *Analysis) GetURL() string                       { return a.url }
func (a *Analysis) GetRequests() int                    { return a.requests }
func (a *Analysis) GetConcurrency() int                 { return a.concurrency }
func (a *Analysis) GetSuccessful() int                   { return a.successful }
func (a *Analysis) GetFailed() int                       { return a.failed }
func (a *Analysis) GetRequestsPerSecond() float64       { return a.requestsPerSecond }
func (a *Analysis) GetFirstByteMin() float64            { return a.firstByteMin }
func (a *Analysis) GetFirstByteMax() float64            { return a.firstByteMax }
func (a *Analysis) GetFirstByteAvg() float64            { return a.firstByteAvg }
func (a *Analysis) GetLastByteMin() float64             { return a.lastByteMin }
func (a *Analysis) GetLastByteMax() float64             { return a.lastByteMax }
func (a *Analysis) GetLastByteAvg() float64             { return a.lastByteAvg }
func (a *Analysis) GetResponseTimeMin() float64         { return a.responseTimeMin }
func (a *Analysis) GetResponseTimeMax() float64         { return a.responseTimeMax }
func (a *Analysis) GetResponseTimeAvg() float64         { return a.responseTimeAvg }
