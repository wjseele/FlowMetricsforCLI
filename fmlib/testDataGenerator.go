package fmlib

import (
	"math/rand"
)

func DataGenerator(datarange, min, max int) ([]int, []int) {
	var totalThroughput int

	var randomThroughputRange = dataGenerator(datarange, min, max)

	for _, n := range randomThroughputRange {
		totalThroughput += randomThroughputRange[n]
	}

	var randomCycletimeRange = dataGenerator(totalThroughput, 1, len(randomThroughputRange))

	return randomThroughputRange, randomCycletimeRange
}

func dataGenerator(datarange, min, max int) []int {
	randomRange := make([]int, datarange)

	for i := 0; i < datarange; i++ {
		randomRange[i] = rand.Intn((max+1)-min) + min
	}

	return randomRange
}
