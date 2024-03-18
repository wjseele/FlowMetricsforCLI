package fmlib

import (
	"testing"
	"time"
)

func TestCycleTimeMax_TwoValuesAreSame_ReturnsZero(t *testing.T) {
	day := time.Date(2030, time.May, 10, 0, 0, 0, 0, time.UTC)
	item := WorkItem{1, "item1", []WorkState{{"start", day}, {"stop", day}}}

	cycleTime := item.CycleTimeMax()

	if cycleTime != 0 {
		t.Errorf("Cycle time was not equal to zero for same time completion. Expected: 0; Actual %d", cycleTime)
	}
}

func TestCycleTimeMax_TwoTimesAreHoursApart_ReturnsOne(t *testing.T) {
	day1 := time.Date(2030, time.May, 10, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.May, 10, 5, 0, 0, 0, time.UTC)
	item := WorkItem{1, "item1", []WorkState{{"start", day1}, {"stop", day2}}}

	cycleTime := item.CycleTimeMax()

	if cycleTime != 1 {
		t.Errorf("Cycle time was not equal to one for same day completion. Expected: 1; Actual %d", cycleTime)
	}
}

func TestCycleTimeMax_TwoDifferentDates_ReturnsDaysDifference(t *testing.T) {
	day1 := time.Date(2030, time.August, 5, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.August, 8, 0, 0, 0, 0, time.UTC)
	item := WorkItem{1, "item1", []WorkState{{"start", day1}, {"stop", day2}}}
	expected := 3

	cycleTime := item.CycleTimeMax()

	if cycleTime != expected {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expected, cycleTime)
	}
}

func TestCycleTimeMax_ManyStates_ReturnsDifferenceBetweenFirstAndLast(t *testing.T) {
	day1 := time.Date(2030, time.March, 29, 10, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.April, 1, 5, 0, 0, 0, time.UTC)
	day3 := time.Date(2030, time.April, 2, 14, 0, 0, 0, time.UTC)
	day4 := time.Date(2030, time.April, 4, 8, 0, 0, 0, time.UTC)

	item := WorkItem{1, "item1", []WorkState{{"state1", day1}, {"state2", day2}, {"state3", day3}, {"state4", day4}}}
	expected := 6

	cycleTime := item.CycleTimeMax()

	if cycleTime != expected {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expected, cycleTime)
	}
}

func TestCycleTimeMax_DatesAreDifferentTimezones_CycleTimeAccountsForTimezones(t *testing.T) {
	day1 := time.Date(2030, time.August, 5, 11, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.August, 11, 11, 0, 0, 0, time.FixedZone("EastlyTime", -7))

	item := WorkItem{1, "item1", []WorkState{{"start", day1}, {"stop", day2}}}
	expected := 7 // Expected because of the timezone offset

	cycleTime := item.CycleTimeMax()

	if cycleTime != expected {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expected, cycleTime)
	}
}

func TestCycleTimeMax_DatesAreReversed_GivesAbsoluteValue(t *testing.T) {
	day1 := time.Date(2030, time.August, 3, 11, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.August, 14, 0, 0, 0, 0, time.UTC)

	item := WorkItem{1, "item1", []WorkState{{"start", day2}, {"stop", day1}}} // Note that day2 comes first
	expected := 11

	cycleTime := item.CycleTimeMax()

	if cycleTime != expected {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expected, cycleTime)
	}
}

func TestCycleTimeMax_LeapYearDates_IncludesFeb29(t *testing.T) {
	day1 := time.Date(2024, time.February, 27, 8, 0, 0, 0, time.UTC)
	day2 := time.Date(2024, time.March, 2, 7, 0, 0, 0, time.UTC)
	item := WorkItem{1, "item1", []WorkState{{"start", day1}, {"stop", day2}}}
	expected := 4

	cycleTime := item.CycleTimeMax()

	if cycleTime != expected {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expected, cycleTime)
	}
}

func TestCycleTime_BetweenTwoStates_ReturnsDaysDifference(t *testing.T) {
	day1 := time.Date(2030, time.March, 29, 10, 0, 0, 0, time.UTC)
	day2 := time.Date(2030, time.April, 1, 5, 0, 0, 0, time.UTC)
	day3 := time.Date(2030, time.April, 2, 14, 0, 0, 0, time.UTC)
	day4 := time.Date(2030, time.April, 4, 8, 0, 0, 0, time.UTC)

	item := WorkItem{1, "item1", []WorkState{{"state1", day1}, {"state2", day2}, {"state3", day3}, {"state4", day4}}}
	expectedCycleTime := 4

	cycleTime := item.CycleTime("state1", "state3")

	if cycleTime != expectedCycleTime {
		t.Errorf("Cycle time mismatch. Expected: %d; Actual: %d", expectedCycleTime, cycleTime)
	}
}

func TestFindWorkState_StateExists_RetrievesIndex(t *testing.T) {
	wsOne := WorkState{"state1", time.Now()}
	wsTwo := WorkState{"state2", time.Now()}
	wsThree := WorkState{"state3", time.Now()}
	wsFour := WorkState{"state4", time.Now()}
	wsFive := WorkState{"state5", time.Now()}

	wi := WorkItem{1, "item1", []WorkState{wsOne, wsTwo, wsThree, wsFour, wsFive}}

	expected := 3 // Zero-based
	actual := wi.findWorkState("state4")

	if expected != actual {
		t.Errorf("Incorrect index identified. Expected: %d; Actual: %d", expected, actual)
	}
}

func TestFindWorkState_StateIsDuplicated_RetrievesFirstInstanceIndex(t *testing.T) {
	wsOne := WorkState{"state1", time.Now()}
	wsTwo := WorkState{"state2", time.Now()}
	wsThree := WorkState{"state5", time.Now()}
	wsFour := WorkState{"state4", time.Now()}
	wsFive := WorkState{"state5", time.Now()}

	wi := WorkItem{1, "item1", []WorkState{wsOne, wsTwo, wsThree, wsFour, wsFive}}

	expected := 2 // Zero-based
	actual := wi.findWorkState("state5")

	if expected != actual {
		t.Errorf("Incorrect index identified. Expected: %d; Actual: %d", expected, actual)
	}
}

func TestFindWorkState_StateDoesNotExist_ReturnsNegativeOne(t *testing.T) {
	wsOne := WorkState{"state1", time.Now()}
	wsTwo := WorkState{"state2", time.Now()}
	wsThree := WorkState{"state3", time.Now()}
	wsFour := WorkState{"state4", time.Now()}
	wsFive := WorkState{"state5", time.Now()}

	wi := WorkItem{1, "item1", []WorkState{wsOne, wsTwo, wsThree, wsFour, wsFive}}

	expected := -1 // Zero-based
	actual := wi.findWorkState("DNE")

	if expected != actual {
		t.Errorf("Incorrect index identified. Expected: %d; Actual: %d", expected, actual)
	}
}
