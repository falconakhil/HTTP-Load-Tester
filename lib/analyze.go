package lib

import (
	"loadtest/models"
	"time"
)

func analyzeResponses(requests int, response chan Response, startTime time.Time, analysis *models.Analysis) {
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
	models.AddData(
		analysis,

		success,
		failure,
		float64(requests)/endTime.Seconds(),

		firstByteMin,
		firstByteMax,
		firstByteAvg/float64(success),

		lastByteMin,
		lastByteMax,
		lastByteAvg/float64(success),

		responseTimeMin,
		responseTimeMax,
		responseTimeAvg/float64(success),
	)
	close(response)
}
