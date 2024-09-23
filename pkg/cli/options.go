package cli

import (
	"github.com/hellflame/argparse"
)

type RunnerOptions struct {
	Endpoint     string
	NumberOfTest int
	Timeout      int
}

func GetOptions() (*RunnerOptions, error) {
	parser := argparse.NewParser("runner", "load test runner for http request", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	endpoint := parser.String("u", "uri", &argparse.Option{
		Required: true,
	})

	numberOfTest := parser.Int("n", "num-request", &argparse.Option{
		Required: true,
	})

	timeout := parser.Int("to", "timeout", &argparse.Option{
		Required: false,
		Default:  "5000",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &RunnerOptions{
		Endpoint:     *endpoint,
		NumberOfTest: *numberOfTest,
		Timeout:      *timeout,
	}, nil
}
