package fmlib

type WorkflowState struct {
	name               string
	workInProcessLimit int
}

func (state WorkflowState) Name() string {
	return state.name
}

func (state WorkflowState) WorkInProcessLimit() int {
	return state.workInProcessLimit
}

func (state WorkflowState) SetWorkInProcessLimit(newLimit int) {
	state.workInProcessLimit = newLimit
}
