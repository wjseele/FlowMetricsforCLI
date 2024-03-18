package fmlib

import (
	"math"
	"time"
)

type WorkState struct {
	stateName string
	date      time.Time
}

type WorkItem struct {
	id     int
	name   string
	states []WorkState
}

func (wi WorkItem) Id() int {
	return wi.id
}

func (wi WorkItem) Name() string {
	return wi.name
}

func (wi WorkItem) States() []WorkState {
	return wi.states
}

func (wi WorkItem) retrieveWorkState(stateName string) WorkState {
	index := wi.findWorkState(stateName)

	return wi.states[index]
}

func (wi WorkItem) findWorkState(stateName string) int {
	for i, item := range wi.states {
		if item.stateName == stateName {
			return i
		}
	}

	return -1
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func cycleTime(first time.Time, last time.Time) int {
	timestampOne := truncateToDay(first)
	timestampTwo := truncateToDay(last)

	diff := timestampTwo.Sub(timestampOne).Abs()
	ct := int(math.Ceil(diff.Hours() / 24))

	if ct == 0 && !first.Equal(last) {
		ct = 1
	}

	return ct
}

func (wi WorkItem) CycleTime(stateOne string, stateTwo string) int {
	first := wi.retrieveWorkState(stateOne)
	second := wi.retrieveWorkState(stateTwo)

	return cycleTime(first.date, second.date)
}

func (wi WorkItem) CycleTimeMax() int {
	first := wi.states[0].date
	last := wi.states[len(wi.states)-1].date

	return cycleTime(first, last)
}
