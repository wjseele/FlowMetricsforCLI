package fmlib

import (
	"math/rand"
	"slices"
)

func HowMany(daysToForecast int, throughputRange []int) (int, int, int, int, []int) {
	var project50, project75, project85, project95 int
	var monteCarloHMSlice = monteCarloHM(daysToForecast, throughputRange)

	total := 10000
	for i := 0; i < len(monteCarloHMSlice); i++ {
		total -= monteCarloHMSlice[i]
		if total >= 5000 {
			project50 = i
		}
		if total >= 7500 {
			project75 = i
		}
		if total >= 8500 {
			project85 = i
		}
		if total >= 9500 {
			project95 = i
		}
	}

	return project50, project75, project85, project95, monteCarloHMSlice
}

func monteCarloHM(daysToForecast int, throughputRange []int) []int {
	var tempTotal int

	monteCarloHMSlice := make([]int, (slices.Max(throughputRange)*daysToForecast)+1)

	for i := 0; i < 10000; i++ {
		tempTotal = monteCarloSingleForecast(daysToForecast, throughputRange)
		monteCarloHMSlice[tempTotal] += 1
	}

	return monteCarloHMSlice
}

func monteCarloSingleForecast(daysToForecast int, throughputRange []int) int {
	total := 0

	for i := 0; i < daysToForecast; i++ {
		total += throughputRange[rand.Intn(len(throughputRange))]
	}

	return total
}
