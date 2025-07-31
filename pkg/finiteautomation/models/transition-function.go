package models

// TransitionFunction represents a single transition function within a finite automaton.
// It contains:
// - currentState: the starting state.
// - input: input string to trigger the transition
// - transitionState: the state where it transitions to
type TransitionFunction struct {
	currentState    *State
	input           string
	transitionState *State
}

func (t *TransitionFunction) Initialize(currentState *State, input string, transitionState *State) {
	t.currentState = currentState
	t.input = input
	t.transitionState = transitionState
}

func (t *TransitionFunction) GetCurrentState() *State {
	return t.currentState
}

func (t *TransitionFunction) GetInput() string {
	return t.input
}

func (t *TransitionFunction) GetTransitionState() *State {
	return t.transitionState
}
