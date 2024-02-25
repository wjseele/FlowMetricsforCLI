package fmlib

import (
	"math/rand"
	"slices"
)

func MonteCarlo(daysToForecast int, throughputRange []int) []int {
	var tempTotal int

	monteCarloSlice := make([]int, (slices.Max(throughputRange)*daysToForecast)+1)

	for i := 0; i < 10000; i++ {
		tempTotal = monteCarlo(daysToForecast, throughputRange)
		monteCarloSlice[tempTotal] += 1
	}

	return monteCarloSlice
}

func monteCarlo(daysToForecast int, throughputRange []int) int {
	total := 0

	for i := 0; i < daysToForecast; i++ {
		total += throughputRange[rand.Intn(len(throughputRange))]
	}

	return total
}
