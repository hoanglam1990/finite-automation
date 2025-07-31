package models

import (
	"errors"
	"fmt"
)

// FiniteAutomation defines a finite automaton model.
// It includes:
// - states: a set of all possible states in the automaton.
// - inputs: the valid input symbols the automaton can process.
// - initialState: the starting state of the automaton.
// - acceptingStates: the set of final states that signify acceptance of input.
// - transitionFunctions: a list of all defined transitions between states.
type FiniteAutomation struct {
	states              map[*State]*State
	inputs              map[string]bool
	initialState        *State
	acceptingStates     map[*State]bool
	transitionFunctions []TransitionFunction
}

// Function to initialize the FiniteAutomation
//   - Check if all attributes are valid
//   - Apply transition function to finite set of states
func (fa *FiniteAutomation) InitializeFiniteAutomation(
	states map[*State]*State,
	inputs map[string]bool,
	initialState *State,
	acceptingStates []*State,
	transitionFunctions []TransitionFunction,
) error {
	if states == nil || initialState == nil || acceptingStates == nil {
		return errors.New(fmt.Sprintln("Invalid nil pointer"))
	}

	err := AreAcceptingStatesValid(states, acceptingStates)
	if err != nil {
		return err
	}

	err = IsInitialStateValid(states, initialState)
	if err != nil {
		return err
	}

	err = AreTransitionFunctionsValid(states, inputs, transitionFunctions)
	if err != nil {
		return err
	}

	fa.states = states
	fa.inputs = inputs
	fa.initialState = initialState
	fa.acceptingStates = map[*State]bool{}
	fa.transitionFunctions = transitionFunctions

	for _, acceptingState := range acceptingStates {
		fa.acceptingStates[acceptingState] = true
	}

	for _, transitionFunction := range fa.transitionFunctions {
		transitionFunction.currentState.addTransition(transitionFunction.input, transitionFunction.transitionState)
	}

	return nil
}

// Function to compute the final state - returns the value of the final state
// - check if there is any nil pointer in the attributes - return error if there is
// - check if the last state is in the list of accepting states - return error if it's not
func (fa *FiniteAutomation) Compute(input string) (*string, error) {
	if fa == nil || fa.states == nil || fa.initialState == nil || fa.acceptingStates == nil {
		return nil, errors.New("finite Automation has not been initialized")
	}

	ref := fa.initialState
	for _, char := range input {
		s := string(char)
		_, isInputValid := fa.inputs[s]
		if !isInputValid {
			return nil, errors.New(fmt.Sprintln("Invalid input: ", s))
		}

		_, isTransitionValid := ref.transition[s]
		if !isTransitionValid {
			return nil, errors.New(fmt.Sprintln("Invalid transition: ", s))
		}

		ref = ref.transition[s]
	}

	if _, valid := fa.acceptingStates[ref]; !valid {
		return nil, errors.New(fmt.Sprintln("Invalid final state - not in the list of accepting state -", ref.output))
	}

	result := ref.GetOutput()

	return &result, nil
}
