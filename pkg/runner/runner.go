package runner

import (
	"log"
	"net/http"
	"runner/pkg/cli"
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

func RunDesiredNumberOfTest(runnerOptions *cli.RunnerOptions) {
	// leaving here until unit test are updated for concurrent req
	count := 1
	request, err := http.NewRequest("GET", runnerOptions.Endpoint, nil)
	if err != nil {
		log.Printf("err: %v \n", err)
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(runnerOptions.Timeout),
	}

	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		_, err := client.Do(request)
		if err != nil {
			log.Printf("err: %v \n", err)
		}

		log.Printf("number of test ran: %v \n", count)
		count++
	}
}
