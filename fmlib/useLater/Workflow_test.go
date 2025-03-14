package fmlib

import "testing"

func TestStates_ReturnsWorkflowStates(t *testing.T) {
	states := []WorkflowState{{"one", 1}, {"two", 1}, {"three", 1}}
	flow := Workflow{states}

	actual := flow.States()

	for i, v := range actual {
		if v.Name() != states[i].Name() || v.WorkInProcessLimit() != states[i].WorkInProcessLimit() {
			t.Errorf("Incorrect states found.  Expected: %v, %d; Actual: %v, %d", v.Name(), v.WorkInProcessLimit(), states[i].Name(), states[i].WorkInProcessLimit())
		}
	}
}
