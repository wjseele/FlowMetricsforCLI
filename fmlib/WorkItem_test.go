package fmlib

import (
	"testing"
	"time"
)

func TestCycleTime_TwoValuesAreSame_ReturnsOne(t *testing.T) {
	day := time.Date(2030, time.May, 10, 0, 0, 0, 0, time.UTC)
	item := WorkItem{1, "item1", []WorkState{{"start", day}, {"stop", day}}}

	cycleTime := item.CycleTime()

	if cycleTime != 1 {
		t.Error("Cycle time was not equal to one for same day completion")
	}
}
