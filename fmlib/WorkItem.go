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

func (wi WorkItem) CycleTime() int {
	first := wi.states[0].date
	last := wi.states[len(wi.states)-1].date

	diff := last.Sub(first).Abs()

	return int(math.Ceil(diff.Hours() / 24))
}
