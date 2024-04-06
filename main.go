package main

import (
	"fmt"

	"github.com/wjseele/FlowMetricsforCLI/fmdemos"
	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI.")
	fmt.Println("There's currently only a demo.")

	selection := 10

	for selection != 0 {
		fmt.Println("Select what you'd like to test.")
		fmt.Println("1: Demo, 2: Parser, 0: Exit")
		_, err := fmt.Scanln(&selection)

		if err != nil {
			fmt.Println(err)
		}

		if selection == 1 {
			fmdemos.SimpleDemo()
		}
		if selection == 2 {
			fmlib.Parser()
		}
	}
}
