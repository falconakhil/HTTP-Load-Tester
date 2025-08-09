package lib

import (
	"fmt"
	"loadtest/models"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)



func PlotGraphs(output_dir string, collection models.AnalysisCollection) error {
    // Create output directory if it doesn't exist
    if err := os.MkdirAll(output_dir, 0755); err != nil {
        return fmt.Errorf("error creating output directory: %v", err)
    }

    // Plot Success/Failure Percentages
    if err := plotSuccessFailurePercentages(output_dir, collection); err != nil {
        return fmt.Errorf("error plotting success/failure: %v", err)
    }

    // Plot Requests Per Second
    if err := plotRequestsPerSecond(output_dir, collection); err != nil {
        return fmt.Errorf("error plotting requests per second: %v", err)
    }

    // Plot Response Times
    if err := plotResponseTimes(output_dir, collection); err != nil {
        return fmt.Errorf("error plotting response times: %v", err)
    }

    // Plot First Byte Times
    if err := plotFirstByteTimes(output_dir, collection); err != nil {
        return fmt.Errorf("error plotting first byte times: %v", err)
    }

    // Plot Last Byte Times
    if err := plotLastByteTimes(output_dir, collection); err != nil {
        return fmt.Errorf("error plotting last byte times: %v", err)
    }

    return nil
}

func plotSuccessFailurePercentages(output_dir string, collection models.AnalysisCollection) error {
    p := plot.New()
    p.Title.Text = "Success/Failure Percentage vs Concurrency"
    p.X.Label.Text = "Concurrency"
    p.Y.Label.Text = "Percentage (%)"

    successPts := make(plotter.XYs, len(collection.Results))
    failurePts := make(plotter.XYs, len(collection.Results))

    for i, result := range collection.Results {
        total := result.GetSuccessful() + result.GetFailed()
        successPercentage := float64(result.GetSuccessful()) / float64(total) * 100
        failurePercentage := float64(result.GetFailed()) / float64(total) * 100

        successPts[i].X = float64(result.GetConcurrency())
        successPts[i].Y = successPercentage

        failurePts[i].X = float64(result.GetConcurrency())
        failurePts[i].Y = failurePercentage
    }

    successLine, err := plotter.NewLine(successPts)
    if err != nil {
        return err
    }
    successLine.Color = plotutil.Color(0)

    failureLine, err := plotter.NewLine(failurePts)
    if err != nil {
        return err
    }
    failureLine.Color = plotutil.Color(1)

    p.Add(successLine, failureLine)
    p.Legend.Add("Success %", successLine)
    p.Legend.Add("Failure %", failureLine)

    return p.Save(8*vg.Inch, 6*vg.Inch, filepath.Join(output_dir, "success_failure_percentage.png"))
}

func plotRequestsPerSecond(output_dir string, collection models.AnalysisCollection) error {
    p := plot.New()
    p.Title.Text = "Requests Per Second vs Concurrency"
    p.X.Label.Text = "Concurrency"
    p.Y.Label.Text = "Requests Per Second"

    pts := make(plotter.XYs, len(collection.Results))
    for i, result := range collection.Results {
        pts[i].X = float64(result.GetConcurrency())
        pts[i].Y = result.GetRequestsPerSecond()
    }

    line, err := plotter.NewLine(pts)
    if err != nil {
        return err
    }
    line.Color = plotutil.Color(2)

    p.Add(line)
    return p.Save(8*vg.Inch, 6*vg.Inch, filepath.Join(output_dir, "requests_per_second.png"))
}

func plotResponseTimes(output_dir string, collection models.AnalysisCollection) error {
    p := plot.New()
    p.Title.Text = "Response Times vs Concurrency"
    p.X.Label.Text = "Concurrency"
    p.Y.Label.Text = "Response Time (ms)"

    minPts := make(plotter.XYs, len(collection.Results))
    maxPts := make(plotter.XYs, len(collection.Results))
    avgPts := make(plotter.XYs, len(collection.Results))

    for i, result := range collection.Results {
        minPts[i].X = float64(result.GetConcurrency())
        minPts[i].Y = result.GetResponseTimeMin()

        maxPts[i].X = float64(result.GetConcurrency())
        maxPts[i].Y = result.GetResponseTimeMax()

        avgPts[i].X = float64(result.GetConcurrency())
        avgPts[i].Y = result.GetResponseTimeAvg()
    }

    minLine, err := plotter.NewLine(minPts)
    if err != nil {
        return err
    }
    minLine.Color = plotutil.Color(0)

    maxLine, err := plotter.NewLine(maxPts)
    if err != nil {
        return err
    }
    maxLine.Color = plotutil.Color(1)

    avgLine, err := plotter.NewLine(avgPts)
    if err != nil {
        return err
    }
    avgLine.Color = plotutil.Color(2)

    p.Add(minLine, maxLine, avgLine)
    p.Legend.Add("Min", minLine)
    p.Legend.Add("Max", maxLine)
    p.Legend.Add("Avg", avgLine)

    return p.Save(8*vg.Inch, 6*vg.Inch, filepath.Join(output_dir, "response_times.png"))
}

func plotFirstByteTimes(output_dir string, collection models.AnalysisCollection) error {
    p := plot.New()
    p.Title.Text = "First Byte Times vs Concurrency"
    p.X.Label.Text = "Concurrency"
    p.Y.Label.Text = "First Byte Time (ms)"

    minPts := make(plotter.XYs, len(collection.Results))
    maxPts := make(plotter.XYs, len(collection.Results))
    avgPts := make(plotter.XYs, len(collection.Results))

    for i, result := range collection.Results {
        minPts[i].X = float64(result.GetConcurrency())
        minPts[i].Y = result.GetFirstByteMin()

        maxPts[i].X = float64(result.GetConcurrency())
        maxPts[i].Y = result.GetFirstByteMax()

        avgPts[i].X = float64(result.GetConcurrency())
        avgPts[i].Y = result.GetFirstByteAvg()
    }

    minLine, err := plotter.NewLine(minPts)
    if err != nil {
        return err
    }
    minLine.Color = plotutil.Color(0)

    maxLine, err := plotter.NewLine(maxPts)
    if err != nil {
        return err
    }
    maxLine.Color = plotutil.Color(1)

    avgLine, err := plotter.NewLine(avgPts)
    if err != nil {
        return err
    }
    avgLine.Color = plotutil.Color(2)

    p.Add(minLine, maxLine, avgLine)
    p.Legend.Add("Min", minLine)
    p.Legend.Add("Max", maxLine)
    p.Legend.Add("Avg", avgLine)

    return p.Save(8*vg.Inch, 6*vg.Inch, filepath.Join(output_dir, "first_byte_times.png"))
}

func plotLastByteTimes(output_dir string, collection models.AnalysisCollection) error {
    p := plot.New()
    p.Title.Text = "Last Byte Times vs Concurrency"
    p.X.Label.Text = "Concurrency"
    p.Y.Label.Text = "Last Byte Time (ms)"

    minPts := make(plotter.XYs, len(collection.Results))
    maxPts := make(plotter.XYs, len(collection.Results))
    avgPts := make(plotter.XYs, len(collection.Results))

    for i, result := range collection.Results {
        minPts[i].X = float64(result.GetConcurrency())
        minPts[i].Y = result.GetLastByteMin()

        maxPts[i].X = float64(result.GetConcurrency())
        maxPts[i].Y = result.GetLastByteMax()

        avgPts[i].X = float64(result.GetConcurrency())
        avgPts[i].Y = result.GetLastByteAvg()
    }

    minLine, err := plotter.NewLine(minPts)
    if err != nil {
        return err
    }
    minLine.Color = plotutil.Color(0)

    maxLine, err := plotter.NewLine(maxPts)
    if err != nil {
        return err
    }
    maxLine.Color = plotutil.Color(1)

    avgLine, err := plotter.NewLine(avgPts)
    if err != nil {
        return err
    }
    avgLine.Color = plotutil.Color(2)

    p.Add(minLine, maxLine, avgLine)
    p.Legend.Add("Min", minLine)
    p.Legend.Add("Max", maxLine)
    p.Legend.Add("Avg", avgLine)

    return p.Save(8*vg.Inch, 6*vg.Inch, filepath.Join(output_dir, "last_byte_times.png"))
}