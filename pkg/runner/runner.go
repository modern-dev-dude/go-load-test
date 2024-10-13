package runner

import (
	"io"
	"log"
	"net/http"
	"runner/pkg/cli"
	"time"

	"github.com/google/uuid"
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
	// leaving here until unit test are updated for concurrent req
	count := 1
	successCount := 0
	errorCount := 0
	// request, err := http.NewRequest("GET", runnerOptions.Endpoint, nil)
	// if err != nil {
	// 	log.Printf("err: %v \n", err)
	// }

	// request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(runnerOptions.Timeout) * time.Second,
	}

	settings := getDefaultSettings()

	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		reqId := uuid.New()

		log.Println("sending req id: ", reqId.String())
		req, err := http.NewRequest(settings.Method, runnerOptions.Endpoint, settings.Body)
		req.Header.Set("Content-Type", settings.ContentType)

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			errorCount++
		}

		defer func() {
			if res != nil {
				log.Println("successful req id: ", reqId.String())

				if res.Status == string(http.StatusOK) {
					successCount++
				}
			}
		}()

		if err != nil {
			log.Printf("err: %v \n", err)
		}

		log.Printf("number of test ran: %v \n", count)
		count++
	}
}
