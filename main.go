package main

import (
	"fmt"
	"os"

	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI.")
	fmt.Println("There's currently only a demo.")

	selection := 0

	fmt.Println("Select what you'd like to test.")
	fmt.Println("1: Demo, 2: Parser, 0: Exit")
	fmt.Scanln(&selection)

	if selection == 1 {
		fmlib.Demo()
	}
	if selection == 2 {
		fmlib.Parser()
	}
	if selection == 0 {
		os.Exit(0)
	}
}
