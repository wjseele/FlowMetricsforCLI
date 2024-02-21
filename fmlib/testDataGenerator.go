package fmlib

import (
	"math/rand"
)

func DataGenerator() []int {
	var daysOfThroughput = 14
	var minThroughputPerDay = 0
	var maxThroughputPerDay = 4
	randomThroughputRange := make([]int, daysOfThroughput)

	for i := 0; i < daysOfThroughput; i++ {
		randomThroughputRange[i] = rand.Intn((maxThroughputPerDay+1)-minThroughputPerDay) + minThroughputPerDay
	}

	return randomThroughputRange
}
