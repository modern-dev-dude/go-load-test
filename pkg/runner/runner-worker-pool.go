package runner

import (
	"log"
	"net/http"
	"runner/pkg/cli"
	"time"

	"github.com/google/uuid"
)

type RunnerResultChannel struct {
	Response      *http.Response
	ExecutionTime time.Duration
	Error         error
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
	if err != nil {
		return &RunnerResultChannel{
			Response:      nil,
			ExecutionTime: timer(time.Now()),
			Error:         err,
		}
	}
	req.Header.Set("Content-Type", settings.ContentType)
	setHeaders(req, &runnerOptions.Headers)

	res, err := client.Do(req)
	if err != nil || res == nil {
		return &RunnerResultChannel{
			Response:      nil,
			ExecutionTime: timer(time.Now()),
			Error:         err,
		}
	}

	if res != nil {
		defer func() {
			res.Body.Close()
		}()
	}

	return &RunnerResultChannel{
		Response:      res,
		ExecutionTime: timer(time.Now()),
		Error:         nil,
	}
}

func setHeaders(r *http.Request, headers *[]string) {
	i := 0
	h := *headers
	for i < len(h)-1 {
		if i+1 > len(h) {
			key := h[i]
			value := h[i+1]
			r.Header.Set(key, value)
			// expect headers be key,value pair
			log.Printf("\nKey: %v, value: %v\n", key, value)
		}
		i += 2
	}
}

func timer(startTime time.Time) time.Duration {
	return time.Since(startTime)
}
