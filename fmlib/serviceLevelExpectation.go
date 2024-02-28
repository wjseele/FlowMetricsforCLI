package fmlib

import (
	"math"
	"slices"
)

func ServiceLevelExpectation(throughputRange []int) (int, int, int, int) {
	var sle50, sle75, sle85, sle95 int

	sle50 = sle(50, throughputRange)
	sle75 = sle(75, throughputRange)
	sle85 = sle(85, throughputRange)
	sle95 = sle(95, throughputRange)

	return sle50, sle75, sle85, sle95
}

func sle(percent float64, throughputRange []int) int {
	var timeOrLess int
	var spotInRange int

	spotInRange = int(math.Round(((float64(len(throughputRange) - 1)) * (percent / 100))))

	slices.Sort(throughputRange)

	timeOrLess = throughputRange[spotInRange]

	return timeOrLess
}
