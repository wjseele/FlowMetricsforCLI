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
