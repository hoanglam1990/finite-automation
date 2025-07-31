package models

import "testing"

// TestAreAcceptingStatesValid_Valid tests that AreAcceptingStatesValid correctly
// identifies accepting states that are present within the defined finite states.
func TestAreAcceptionStatesValid_Valid(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state1}

	err := AreAcceptingStatesValid(finiteStates, acceptingStates)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// TestAreAcceptingStatesValid_Invalid ensures that AreAcceptingStatesValid returns an error
// when the accepting states include a state not found in the finite state set.
func TestAreAcceptionStatesValid_Invalid(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{})

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state3}

	err := AreAcceptingStatesValid(finiteStates, acceptingStates)

	if err == nil {
		t.Errorf("Expected error for invalid accepting state, got nil")
	}
}

// TestIsInitialStateValid_Valid tests that IsInitialStateValid succeeds
// when the provided initial state is included in the finite state set.
func TestIsInitialStateValid_Valid(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	err := IsInitialStateValid(finiteStates, &state1)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// TestIsInitialStateValid_Invalid tests an invalid initial state 
// that is not part of the finite state set.
func TestIsInitialStateValid_Invalid(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{})

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	err := IsInitialStateValid(finiteStates, &state3) // <-- state3 not in the list

	if err == nil {
		t.Errorf("Expected error for invalid initial state, got nil")
	}
}

// TestAreTransitionFunctionsValid_Valid tests that AreTransitionFunctionsValid correctly
// validates a set of transition functions
func TestAreTransitionFunctionsValid_Valid(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	inputs := map[string]bool{"0": true, "1": true}

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state1)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	err := AreTransitionFunctionsValid(finiteStates, inputs, transitionFunctions)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// TestAreTransitionFunctionsValid_InvalidCurrentState tests that AreTransitionFunctionsValid
// correctly identifies a transition function with an invalid current state.
func TestAreTransitionFunctionsValid_InvalidCurrentState(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{}) // <-- invalid current state
	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	inputs := map[string]bool{"0": true, "1": true}

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state1)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state3, "1", &state1) // <-- transition with invalid starting state (state3 not in the list)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	err := AreTransitionFunctionsValid(finiteStates, inputs, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for transition function, got nil")
	}
}

// TestAreTransitionFunctionsValid_InvalidCurrentState tests that AreTransitionFunctionsValid
// correctly identifies a transition function with an invalid transition state.
func TestAreTransitionFunctionsValid_InvalidTransitionState(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{}) // <-- invalid transition state
	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	inputs := map[string]bool{"0": true, "1": true}

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state1)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state3) // <-- transition with invalid transition state (state3 not in the list)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	err := AreTransitionFunctionsValid(finiteStates, inputs, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for transition function, got nil")
	}
}

// TestAreTransitionFunctionsValid_InvalidCurrentState tests that AreTransitionFunctionsValid
// correctly identifies an input not in the list of finite inputs.
func TestAreTransitionFunctionsValid_InvalidInputs(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}

	inputs := map[string]bool{"0": true, "2": true} // <-- accepting inputs 0 & 2

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state1)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1) // <-- invalid input 1
	transitionFunctions := []TransitionFunction{tf1, tf2}

	err := AreTransitionFunctionsValid(finiteStates, inputs, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for transition function, got nil")
	}
}
