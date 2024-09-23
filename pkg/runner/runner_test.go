package runner_test

import (
	"bytes"
	"log"
	"runner/pkg/cli"
	"runner/pkg/runner"
	"strings"
	"testing"
)

func getOpts() *cli.RunnerOptions {
	return &cli.RunnerOptions{
		Endpoint:     "http://test.com",
		NumberOfTest: 3,
	}
}

// TODO make this a better test when I come back to concurrent test execution
func TestRunDesiredNumberOfTest(t *testing.T) {
	var buff bytes.Buffer
	log.SetOutput(&buff)

	// restores the log after unit test
	defer log.SetOutput(nil)

	runnerOpts := getOpts()
	runner.RunDesiredNumberOfTest(runnerOpts)

	buffAsString := buff.String()

	if strings.Contains(buffAsString, "1") == false {
		t.Errorf("number of desired runs are not correct")
	}

	if strings.Contains(buffAsString, "2") == false {
		t.Errorf("number of desired runs are not correct")
	}

	if strings.Contains(buffAsString, "3") == false {
		t.Errorf("number of desired runs are not correct")
	}

}
