package main

import (
	"math/rand"
	"slices"
	"fmt"
)

func main() {
	var daysOfThroughput int // daysToForecast int
	//var tempTotal, tempRolls int
	var minThroughputPerDay, maxThroughputPerDay int

	daysOfThroughput = 14
	minThroughputPerDay = 0
	maxThroughputPerDay = 4
	randomThroughputRange := make([]int, daysOfThrougput)

	for i := 0; i < daysOfThroughput; i++ {
		randomThroughputRange[i] = rand.Intn(maxThrougputPerDay - minThroughputPerDay) + minThrougputPerDay
	}

	fmt.Println("Throughput for each day is:")

	for i, n := range randomThroughputRange {
		fmt.Println("Day", i, ":", n)
	}
}
