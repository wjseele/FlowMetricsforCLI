package main

import (
	"fmt"
	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI.")
	fmt.Println("Let's try a test.")

	var randomThroughputRange = fmlib.DataGenerator()

	fmt.Println("Throughput for each day is:")
	for i, n := range randomThroughputRange {
		fmt.Println("Day", i+1, ":", n)
	}

	fmt.Println("Now let's try running a Monte Carlo simulation.")
	fmt.Println("We'll forecast 30 days of work on our test data.")

	daysToForecast := 30

	var monteCarloSlice = fmlib.MonteCarlo(daysToForecast, randomThroughputRange)

	fmt.Println("Our forecast is:")
	for i, n := range monteCarloSlice {
		fmt.Println("Amount of", i, ":", n)
	}
}
