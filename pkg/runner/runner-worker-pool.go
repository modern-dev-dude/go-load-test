package runner

import (
	"log"
	"net/http"
	"runner/pkg/cli"
	"time"

	"github.com/google/uuid"
)

type RunnerResultChannel struct {
	Response *http.Response
	Error    error
}

// basic worker pool adapted from https://blog.stackademic.com/5-go-concurrency-patterns-i-wish-i-learned-earlier-bbfc02afc44b
func WorkerPool(numWorkers int, jobQueue <-chan *cli.RunnerOptions, resultQueue chan<- *RunnerResultChannel) {
	for i := 0; i < numWorkers; i++ {
		go worker(jobQueue, resultQueue)
	}
}

func worker(jobQueue <-chan *cli.RunnerOptions, resultQueue chan<- *RunnerResultChannel) {
	for job := range jobQueue {
		resultQueue <- processApiRequest(job)
	}
}

func processApiRequest(runnerOptions *cli.RunnerOptions) *RunnerResultChannel {

	client := &http.Client{
		Timeout: time.Duration(runnerOptions.Timeout) * time.Second,
	}

	settings := getDefaultSettings()

	reqId := uuid.New()

	log.Println("sending req id: ", reqId.String())
	req, err := http.NewRequest(settings.Method, runnerOptions.Endpoint, settings.Body)
	req.Header.Set("Content-Type", settings.ContentType)

	res, err := client.Do(req)
	if err != nil {
		return &RunnerResultChannel{
			Response: nil,
			Error:    err,
		}
	}

	if res != nil {
		defer func() {
			res.Body.Close()
		}()
	}

	return &RunnerResultChannel{
		Response: res,
		Error:    nil,
	}
}
