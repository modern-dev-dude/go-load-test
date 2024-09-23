package runner

import (
	"log"
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

func RunDesiredNumberOfTest(runnerOptions *cli.RunnerOptions) {
	count := 1

	for i := 0; i < runnerOptions.NumberOfTest; i++ {
		log.Printf("Num test run %v\n", count)
		count++
	}
}
