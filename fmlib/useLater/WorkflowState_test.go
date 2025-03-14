package fmlib

import "testing"

func TestName_RetrievesName(t *testing.T) {
	stateName := "This is MY State!"
	wipLimit := 0
	state := WorkflowState{stateName, wipLimit}

	actual := state.Name()

	if actual != stateName {
		t.Errorf("Incorrect name returned. Expected: %v; Actual: %v", stateName, actual)
	}
}

func TestWorkInProcessLimit_RetrievesCurrentLimit(t *testing.T) {
	stateName := "staaate"
	wipLimit := 3
	state := WorkflowState{stateName, wipLimit}

	actual := state.WorkInProcessLimit()

	if actual != wipLimit {
		t.Errorf("Incorrect WIP Limit Returned. Expected: %v; Actual: %v", wipLimit, actual)
	}
}

func TestSetWorkInProcessLimit_ChangeLimit_NewLimitReturned(t *testing.T) {
	stateName := "staaate"
	wipLimit := 3
	newWipLimit := 1
	state := WorkflowState{stateName, wipLimit}
	state.SetWorkInProcessLimit(newWipLimit)

	actual := state.WorkInProcessLimit()

	if actual != wipLimit {
		t.Errorf("Incorrect WIP Limit Returned. Expected: %v; Actual: %v", newWipLimit, actual)
	}
}
