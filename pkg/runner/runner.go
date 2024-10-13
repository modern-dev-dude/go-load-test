package runner

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runner/pkg/cli"
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
	Method      string
	ContentType string
	Body        io.Reader
}

func getDefaultSettings() *RunnerSettings {
	return &RunnerSettings{
		Method:      "GET",
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

	// collect results
	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		result := <-results
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

	fmt.Printf("\nSuccesses :%v\n", successCount)
	fmt.Printf("\nErrors :%v\n", errorCount)
}
