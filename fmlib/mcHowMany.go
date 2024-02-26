package fmlib

func HowMany(daysToForecast int, throughputRange []int) (int, int, int, int, []int) {
	var project50, project75, project85, project95 int
	var monteCarloSlice = MonteCarlo(daysToForecast, throughputRange)

	total := 10000
	for i := 0; i < len(monteCarloSlice); i++ {
		total -= monteCarloSlice[i]
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

	return project50, project75, project85, project95, monteCarloSlice
}
