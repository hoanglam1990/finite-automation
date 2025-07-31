package models

// State represents a single state within a finite automaton.
// It contains:
// - output: the associated output value of this state.
// - transition: a mapping of input symbols (keys) to the next state,
//   	allowing traversal through the automaton based on input.
type State struct {
	output     string
	transition map[string]*State
}

func (st *State) Initialize(output string, transition map[string]*State) {
	st.output = output
	st.transition = transition
}

func (st *State) GetOutput() string {
	return st.output
}

// addTransition adds or updates a transition for the given input symbol.
// 	- If a transition for the input already exists, it will be overwritten
// 	with the new target state.
func (st *State) addTransition(input string, state *State) {
	st.transition[input] = state
}
