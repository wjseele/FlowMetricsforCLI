package fmlib

import "fmt"

type item struct {
	ID       int64
	name     string
	created  int
	state1   int
	state2   int
	state3   int
	finished int
	blocked  bool
}

func Parser() {
	wobbles := item{1, "Wobbles", 1, 2, 3, 4, 5, true}
	fmt.Println(wobbles)
	fmt.Println(wobbles.name)
	fmt.Println(wobbles.finished)
}

func newParser() map[string]map[string]string {
	//Create a 2 dimensional map containing item history
	//Read each line from the database into a map
	//ID - Name - StateDate - BlockedDays - Tags

	flowMetrics := make(map[string]map[string]string)

	return flowMetrics
}
