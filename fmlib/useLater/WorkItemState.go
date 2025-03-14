package fmlib

import (
	"time"
)

type WorkItemState struct {
	stateName string
	date      time.Time
}
