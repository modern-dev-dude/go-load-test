package cli

import (
	"github.com/hellflame/argparse"
)

type RunnerOptions struct {
	Endpoint     string
	NumberOfTest int
	Timeout      int
	Connections  int
	Headers      []string
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
		Default:  "5",
	})

	connections := parser.Int("c", "connections", &argparse.Option{
		Required: false,
		Default:  "5",
	})

	headers := parser.Strings("he", "headers", &argparse.Option{
		Required: false,
	})

	if headers == nil {
		headers = &[]string{}
	}

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &RunnerOptions{
		Endpoint:     *endpoint,
		NumberOfTest: *numberOfTest,
		Timeout:      *timeout,
		Connections:  *connections,
		Headers:      *headers,
	}, nil
}
