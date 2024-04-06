package fmdemos

import (
	"fmt"

	"github.com/wjseele/FlowMetricsforCLI/fmlib"
)

func SimpleDemo() {
	var days, min, max int

	fmt.Println("Enter the values you wish to use.")
	fmt.Println("How many historical days of work you wish to generate?")
	fmt.Scanln(&days)
	fmt.Println("What's the minimum throughput achieved per day?")
	fmt.Scanln(&min)
	fmt.Println("What's the maximum throughput achieved per day?")
	fmt.Scanln(&max)

	fmt.Println("Gotcha, I'll simulate", days, "days of work, with a minimum of", min, "items per day and a maximum of", max)
	fmt.Println("I'll even generate some cycle times for those items.")
	fmt.Println("Here we go!")

	var randomThroughputRange, randomCycletimeRange = DataGenerator(days, min, max)

	exit := false
	selection := 0

	for exit != true {
		fmt.Println("Select demo function:")
		fmt.Println("1. How much can we get done?")
		fmt.Println("2. How long until I get a certain amount of items?")
		fmt.Println("3. What's the Service Level Expectation?")
		fmt.Println("4. Exit.")
		fmt.Scanln(&selection)
		if selection == 1 {
			howMuchDemo(randomThroughputRange)
		}
		if selection == 2 {
			howLongDemo(randomThroughputRange)
		}
		if selection == 3 {
			sleDemo(randomCycletimeRange)
		}
		if selection == 4 {
			exit = true
		}
	}
}

func howMuchDemo(throughputRange []int) {
	var daysToForecast int

	fmt.Println("Now let's try running a Monte Carlo simulation.")
	fmt.Println("How many days do you want to forecast?")
	fmt.Scanln(&daysToForecast)

	var howmuch50, howmuch75, howmuch85, howmuch95, monteCarloSlice = fmlib.HowMany(daysToForecast, throughputRange)

	fmt.Println("Our forecast is:")
	for i, n := range monteCarloSlice {
		if n != 0 {
			fmt.Println("Times we finished", i, "items:", n)
		}
	}

	fmt.Println("Now let's see how much we could get done.")
	fmt.Println("There's a 50% chance of delivering at least", howmuch50, "items.")
	fmt.Println("There's a 75% chance of delivering at least", howmuch75, "items.")
	fmt.Println("There's a 85% chance of delivering at least", howmuch85, "items.")
	fmt.Println("There's a 95% chance of delivering at least", howmuch95, "items.")
}

func howLongDemo(throughputRange []int) {
	var desiredAmount int

	fmt.Println("Now let's see how long it takes us to deliver a certain amount of items.")
	fmt.Println("How many items do you want to deliver?")
	fmt.Scanln(&desiredAmount)

	var howlong50, howlong75, howlong85, howlong95, monteCarloHLSlice = fmlib.HowLong(desiredAmount, throughputRange)

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
}

func sleDemo(cycletimeRange []int) {

	fmt.Println("Let's calculate the SLE's for this range of throughputs.")

	var sle50, sle75, sle85, sle95 = fmlib.ServiceLevelExpectation(cycletimeRange)

	fmt.Println("In this range, we historically delivered...")
	fmt.Println("50 % of items in", sle50, "days or less.")
	fmt.Println("75 % of items in", sle75, "days or less.")
	fmt.Println("85 % of items in", sle85, "days or less.")
	fmt.Println("95 % of items in", sle95, "days or less.")
}
