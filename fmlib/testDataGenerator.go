package fmlib

import (
	//	"fmt"
	"math/rand"
	//"slices"
)

func DataGenerator() []int {
	// var daysOfThroughput, daysToForecast int
	// var tempTotal, tempRolls int
	// var minThroughputPerDay, maxThroughputPerDay int

	var daysOfThroughput = 14
	var minThroughputPerDay = 0
	var maxThroughputPerDay = 4
	randomThroughputRange := make([]int, daysOfThroughput)

	for i := 0; i < daysOfThroughput; i++ {
		randomThroughputRange[i] = rand.Intn((maxThroughputPerDay+1)-minThroughputPerDay) + minThroughputPerDay
	}

	return randomThroughputRange

	//fmt.Println("Throughput for each day is:")

	//for i, n := range randomThroughputRange {
	//	fmt.Println("Day", i+1, ":", n)
	//}
}
