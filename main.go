package main

import (
	"fmt"
	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI.")
	fmt.Println("There's currently only a demo.")

	fmlib.Demo()
	fmlib.Parser()
}
