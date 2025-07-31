package models

import (
	"errors"
	"fmt"
)

// Function to check if the accepting states F is the subset of finite states Q
// - returns error if any accepting state not found in Q
func AreAcceptingStatesValid(states map[*State]*State, acceptingStates []*State) error {
	for _, state := range acceptingStates {
		if _, ok := states[state]; !ok {
			return errors.New(fmt.Sprintln("Accepting State invalid - Accepting state not in the set of states: ", state.GetOutput()))
		}
	}

	return nil
}

// Function to check if the initial state q0 is in the set of finite states Q
func IsInitialStateValid(states map[*State]*State, initialState *State) error {
	if _, ok := states[initialState]; !ok {
		return errors.New(fmt.Sprintln("Initial State invalid - Initial state not in the set of states: ", initialState.GetOutput()))
	}

	return nil
}

// Function to check if the transitions are valid
//   - both starting state and transition state must be in the sets of finite states
//   - input symbol must be in the set of finite inputs
func AreTransitionFunctionsValid(
	states map[*State]*State,
	inputs map[string]bool,
	transitionFunctions []TransitionFunction,
) error {
	for _, transitionFunction := range transitionFunctions {
		if _, ok := states[transitionFunction.GetCurrentState()]; !ok {
			return errors.New(fmt.Sprintln("Transition Function invalid - Starting state not in the set of states: ", transitionFunction.GetCurrentState().GetOutput()))
		}

		if _, ok := states[transitionFunction.GetTransitionState()]; !ok {
			return errors.New(fmt.Sprintln("Transition Function invalid - Transition state not in the set of states: ", transitionFunction.GetTransitionState().GetOutput()))
		}

		if _, ok := inputs[transitionFunction.GetInput()]; !ok {
			return errors.New(fmt.Sprintln("Transition Function invalid - Input not in the set of finite inputs", transitionFunction.GetInput()))
		}
	}

	return nil
}
