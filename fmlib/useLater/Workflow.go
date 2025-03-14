package fmlib

type Workflow struct {
	states []WorkflowState
}

func (flow Workflow) States() []WorkflowState {
	return flow.states
}
