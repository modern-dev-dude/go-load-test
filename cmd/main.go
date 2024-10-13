package main

import (
	"fmt"
	"runner/pkg/runner"
	"time"
)

func main() {
	timeTotalExecution := time.Now()

	runner.Start()

	fmt.Printf("\nDone success: %v\n", time.Since(timeTotalExecution))
}
