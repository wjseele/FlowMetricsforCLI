package fmlib

import (
	"math/rand"
)

func DataGenerator(datarange, min, max int) []int {
	randomRange := make([]int, datarange)

	for i := 0; i < datarange; i++ {
		randomRange[i] = rand.Intn((max+1)-min) + min
	}

	return randomRange
}
