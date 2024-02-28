package fmlib

import (
	"fmt"
	"math/rand"
	"slices"
)

func HowLong(desiredAmount int, throughputRange []int) (int, int, int, int, []int) {
	var project50, project75, project85, project95 int

	var monteCarloHLSlice = howLong(desiredAmount, throughputRange)

	total := 10000
	for i := len(monteCarloHLSlice) - 1; i >= 0; i-- {
		total -= monteCarloHLSlice[i]
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

	return project50, project75, project85, project95, monteCarloHLSlice
}

func howLong(desiredAmount int, throughputRange []int) []int {
	var minimumDelivery, daysToForecast int

	if slices.Min(throughputRange) > 0 {
		minimumDelivery = slices.Min(throughputRange)
	} else {
		minimumDelivery = 1
	}

	daysToForecast = (desiredAmount / minimumDelivery) * counts(throughputRange, slices.Min(throughputRange))

	var monteCarloHLSlice = monteCarloHL(desiredAmount, daysToForecast, throughputRange)

	return monteCarloHLSlice
}

func counts[S []E, E comparable](s S, e E) int {
	var n int
	for _, v := range s {
		if v == e {
			n++
		}
	}
	return n
}

func monteCarloHL(desiredAmount int, daysToForecast int, throughputRange []int) []int {
	var tempTotal int

	monteCarloHLSlice := make([]int, daysToForecast)

	for i := 0; i < 10000; i++ {
		tempTotal = monteCarloSingleCountdown(desiredAmount, daysToForecast, throughputRange)
		monteCarloHLSlice[tempTotal] += 1
	}

	return monteCarloHLSlice
}

func monteCarloSingleCountdown(desiredAmount int, daysToForecast int, throughputRange []int) int {
	days := 0

	for desiredAmount > 0 && days < daysToForecast {
		desiredAmount -= throughputRange[rand.Intn(len(throughputRange))]
		days++
	}

	if desiredAmount > 0 {
		fmt.Println("Wasn't able to deliver all items within", daysToForecast)
		fmt.Println(desiredAmount, "items left. Please run again with more time.")
	}
	return days
}
