package fmlib

import "time"

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
	return 1
}
