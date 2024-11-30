package runner

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runner/pkg/cli"
	"sort"
	"time"
)

func Start() {
	runnerOptions, err := cli.GetOptions()
	if err != nil {
		// TODO will update the error later
		log.Printf("Error getting args %v\n", err)
		return
	}
	RunDesiredNumberOfTest(runnerOptions)
}

type RunnerSettings struct {
	ContentType string
	Body        io.Reader
}

func getDefaultSettings() *RunnerSettings {
	return &RunnerSettings{
		ContentType: "application/json",
		Body:        nil,
	}
}

func RunDesiredNumberOfTest(runnerOptions *cli.RunnerOptions) {
	jobs := make(chan *cli.RunnerOptions, runnerOptions.NumberOfTest)
	results := make(chan *RunnerResultChannel, runnerOptions.NumberOfTest)
	fmt.Printf("runnerOptions: %+v\n", runnerOptions)
	WorkerPool(runnerOptions.Connections, jobs, results)

	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		jobs <- runnerOptions
	}
	close(jobs)

	errorCount := 0
	successCount := 0
	responseTime := make([]time.Duration, 0)

	// collect results
	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		result := <-results

		// collect response time
		responseTime = append(responseTime, result.ExecutionTime)
		// handle errors
		if result.Error != nil {
			errorCount++
		}

		if result.Response != nil {
			if result.Response.StatusCode == http.StatusOK {
				successCount++
			}
		}
	}

	sort.Slice(responseTime, func(i, j int) bool {
		return int(responseTime[i].Nanoseconds()) > int(responseTime[j].Nanoseconds())
	})

	var totalDuration int64 = 0
	for _, t := range responseTime {
		totalDuration += t.Nanoseconds()
	}

	avg := float64(totalDuration) / float64(len(responseTime))
	pNinetyIdx := int(float32(len(responseTime)) * 0.9)

	fmt.Printf("\nSuccesses :%v\n", successCount)
	fmt.Printf("\nErrors :%v\n", errorCount)
	fmt.Printf("\naverage response time :%v\n", avg)
	fmt.Printf("\np90 response time :%v\n", responseTime[pNinetyIdx])
}
