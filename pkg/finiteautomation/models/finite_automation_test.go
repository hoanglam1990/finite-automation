package models

import (
	"strings"
	"testing"
)

// TestInitializeFiniteAutomation_NoError validates that the initialization
// of a FiniteAutomation with well-defined states, transitions, inputs,
// and accepting conditions completes without error.
func TestInitializeFiniteAutomation_NoError(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state1}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state1

	fa := FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(finiteStates, inputs, initialStates, acceptingStates, transitionFunctions)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// TestInitializeFiniteAutomation_NilPointerError verifies that initializing
// a FiniteAutomation with a nil finiteStates map triggers an error.
func TestInitializeFiniteAutomation_NilPointerError(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	acceptingStates := []*State{&state1}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state1

	fa := FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(nil, inputs, initialStates, acceptingStates, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for invalid nil pointer, got nil")
	}
}

// TestInitializeFiniteAutomation_InvalidAcceptingState tests that the
// initialization fails when the set of accepting states includes a state
// that is not part of the defined finite states.
func TestInitializeFiniteAutomation_InvalidAcceptingState(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state3}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state1

	fa := FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(finiteStates, inputs, initialStates, acceptingStates, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for invalid accepting state, got nil")
	}
}

// TestInitializeFiniteAutomation_InvalidInitialState tests that the initialization
// when the initial state is not included in the set of defined finite states.
func TestInitializeFiniteAutomation_InvalidInitialState(t *testing.T) {
	state1, state2, state3 := State{}, State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})
	state3.Initialize("2", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state1}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state3

	fa := FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(finiteStates, inputs, initialStates, acceptingStates, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for invalid accepting state, got nil")
	}
}

// TestInitializeFiniteAutomation_InvalidTransitionFunction tests that the
// initialization fails when a transition function contains an input symbol
// that is not defined in the automaton’s allowed inputs.
func TestInitializeFiniteAutomation_InvalidTransitionFunction(t *testing.T) {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "2", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state1}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state1

	fa := FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(finiteStates, inputs, initialStates, acceptingStates, transitionFunctions)

	if err == nil {
		t.Errorf("Expected error for invalid transition function, got nil")
	}
}

func GetMockFiniteAutomation() FiniteAutomation {
	state1, state2 := State{}, State{}
	state1.Initialize("0", map[string]*State{})
	state2.Initialize("1", map[string]*State{})

	tf1 := TransitionFunction{}
	tf1.Initialize(&state1, "0", &state2)
	tf2 := TransitionFunction{}
	tf2.Initialize(&state2, "1", &state1)
	transitionFunctions := []TransitionFunction{tf1, tf2}

	finiteStates := map[*State]*State{&state1: &state1, &state2: &state2}
	acceptingStates := []*State{&state1}
	inputs := map[string]bool{"0": true, "1": true}
	initialStates := &state1

	fa := FiniteAutomation{}
	fa.InitializeFiniteAutomation(finiteStates, inputs, initialStates, acceptingStates, transitionFunctions)

	return fa
}

// TestCompute_NoError tests that the Compute function processes a valid input
// string correctly through the finite automaton.
func TestCompute_NoError(t *testing.T) {
	fa := GetMockFiniteAutomation()

	result, err := fa.Compute("01")

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if *result != "0" {
		t.Errorf("Expected %s error, got %s", "0", *result)
	}
}

// TestCompute_ErrorInvalidInput tests that the Compute function returns an error
// when the input string contains symbols not defined in the automaton's input set.
func TestCompute_ErrorInvalidInput(t *testing.T) {
	fa := GetMockFiniteAutomation()

	result, err := fa.Compute("21")

	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}

	if !strings.Contains(err.Error(), "Invalid input") {
		t.Errorf("Expected error %s, got %s", "Invalid input", err.Error())
	}

	if result != nil {
		t.Errorf("Expected nil result, got %s", *result)
	}
}

// TestCompute_ErrorInvalidTransition tests that the Compute function returns
// an error when a transition for a given input does not exist from the current state.
func TestCompute_ErrorInvalidTransition(t *testing.T) {
	fa := GetMockFiniteAutomation()

	result, err := fa.Compute("00")

	if err == nil {
		t.Errorf("Expected error for invalid transition, got nil")
	}

	if !strings.Contains(err.Error(), "Invalid transition") {
		t.Errorf("Expected error %s, got %s", "Invalid input", err.Error())
	}

	if result != nil {
		t.Errorf("Expected nil result, got %s", *result)
	}
}

// TestCompute_ErrorNotInitialized entestsures that calling Compute returns
// an error when the FiniteAutomation fa is nil.
func TestCompute_ErrorNotInitialized(t *testing.T) {
	var fa *FiniteAutomation = nil

	result, err := fa.Compute("10")

	if err == nil {
		t.Errorf("Expected error for invalid transition, got nil")
	}

	if !strings.Contains(err.Error(), "initialized") {
		t.Errorf("Expected error %s, got %s", "Invalid input", err.Error())
	}

	if result != nil {
		t.Errorf("Expected nil result, got %s", *result)
	}
}

// TestCompute_ErrorInvalidFinalState test that the Compute function returns
// an error when the final state reached after processing the input string
// is not an accepting state.
func TestComput_ErrorInvalidFinalState(t *testing.T) {
	fa := GetMockFiniteAutomation()

	result, err := fa.Compute("0")

	if err == nil {
		t.Errorf("Expected error for invalid final state, got nil")
	}

	if !strings.Contains(err.Error(), "Invalid final state") {
		t.Errorf("Expected error %s, got %s", "Invalid input", err.Error())
	}

	if result != nil {
		t.Errorf("Expected nil result, got %s", *result)
	}
}
