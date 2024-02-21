package main

import (
	"fmt"
	"github.com/wjseele/FlowMetricsforCLI/lib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI")
	fmt.Println("Let's try a test")

	var randomThroughputRange = fmlib.DataGenerator()

	fmt.Println("Throughput for each day is:")
	for i, n := range randomThroughputRange {
		fmt.Println("Day", i+1, ":", n)
	}
}
