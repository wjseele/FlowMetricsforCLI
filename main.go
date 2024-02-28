package main

import (
	"fmt"
	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func main() {
	fmt.Println("Welcome to Flow Metrics for CLI.")
	fmt.Println("Let's try a test.")

	var randomThroughputRange, randomCycletimeRange = fmlib.DataGenerator(14, 0, 8)

	fmt.Println("Throughput for each day is:")
	for i, n := range randomThroughputRange {
		fmt.Println("Day", i+1, ":", n)
	}

	fmt.Println("Now let's try running a Monte Carlo simulation.")
	fmt.Println("We'll forecast 30 days of work on our test data.")

	daysToForecast := 30

	var project50, project75, project85, project95, monteCarloSlice = fmlib.HowMany(daysToForecast, randomThroughputRange)

	fmt.Println("Our forecast is:")
	for i, n := range monteCarloSlice {
		if n != 0 {
			fmt.Println("Amount of", i, ":", n)
		}
	}

	fmt.Println("Now let's see how much we could get done.")
	fmt.Println("There's a 50% chance of delivering at least", project50, "items.")
	fmt.Println("There's a 75% chance of delivering at least", project75, "items.")
	fmt.Println("There's a 85% chance of delivering at least", project85, "items.")
	fmt.Println("There's a 95% chance of delivering at least", project95, "items.")

	fmt.Println("Now let's see how long it takes us to deliver a certain amount of items.")
	fmt.Println("We'll aim for 60.")

	desiredAmount := 60

	var howlong50, howlong75, howlong85, howlong95, monteCarloHLSlice = fmlib.HowLong(desiredAmount, randomThroughputRange)

	fmt.Println("Our forecast is:")
	for i, n := range monteCarloHLSlice {
		if n != 0 {
			fmt.Println("It took", i, "days in", n, "cases.")
		}
	}

	fmt.Println("Now let's see how many days we should expect.")
	fmt.Println("There's a 50% chance to deliver", desiredAmount, "in", howlong50, "days.")
	fmt.Println("There's a 75% chance to deliver", desiredAmount, "in", howlong75, "days.")
	fmt.Println("There's a 85% chance to deliver", desiredAmount, "in", howlong85, "days.")
	fmt.Println("There's a 95% chance to deliver", desiredAmount, "in", howlong95, "days.")

	fmt.Println("Now let's calculate the SLE's for this range of throughputs.")

	var sle50, sle75, sle85, sle95 = fmlib.ServiceLevelExpectation(randomCycletimeRange)

	fmt.Println("In this range, we historically delivered...")
	fmt.Println("50 % of items in", sle50, "days or less.")
	fmt.Println("75 % of items in", sle75, "days or less.")
	fmt.Println("85 % of items in", sle85, "days or less.")
	fmt.Println("95 % of items in", sle95, "days or less.")
}
