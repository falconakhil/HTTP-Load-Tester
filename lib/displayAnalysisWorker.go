package lib

import (
	"fmt"
	"loadtest/models"
)

// Display analysis until nil is received
func displayAnalysisWorker(analysis chan *models.Analysis) {
	for a := range analysis {
		if a == nil {
			break
		}
		fmt.Println()
		models.DisplayAnalysis(a)
	}
}
