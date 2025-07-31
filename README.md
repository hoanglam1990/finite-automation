# Finite Automation (Go)

This repository contains an implementation of a finite automation in Go. It includes support for defining states, transitions, accepting conditions, and simulating input processing.

## 🧠 Overview

The automation is composed of:

- `State`: Represents a node in the automation graph with an output and transition map.
- `TransitionFunction`: Defines a rule for moving between states based on an input symbol.
- `FiniteAutomation`: Holds the complete automation, including its states, inputs, initial state, accepting states, and transition logic.

## 🔧 Features

- Create and connect states with transitions.
- Define accepted input symbols.
- Validate structure before simulation.
- Simulate input strings and determine acceptance.
- Comprehensive test suite for validation logic.

## Unit tests

- TestInitializeFiniteAutomation_NoError - Validates successful initialization with correct setup.
- TestInitializeFiniteAutomation_NilPointerError - Verifies error when finiteStates is nil.
- TestInitializeFiniteAutomation_InvalidAcceptingState - Checks error when accepting state is not part of finiteStates.
- TestInitializeFiniteAutomation_InvalidInitialState - Tests error when initial state isn't part of finiteStates.
- TestInitializeFiniteAutomation_InvalidTransitionFunction - Verifies error for transition input not in allowed input set.
- TestCompute_NoError - Validates correct computation of valid input string.
- TestCompute_ErrorInvalidInput - Ensures error is raised for undefined input symbols.
- TestCompute_ErrorInvalidTransition - Checks error when a valid input lacks a defined transition.
- TestCompute_ErrorNotInitialized - Ensures error when FiniteAutomation is not initialized (nil pointer).
- TestComput_ErrorInvalidFinalState - Verifies error when final state is not an accepting state.

- TestAreAcceptionStatesValid_Valid - Verifies that valid accepting states pass validation.
- TestAreAcceptionStatesValid_Invalid - Checks error when accepting states include undefined states.
- TestIsInitialStateValid_Valid - Ensures initial state is correctly recognized within finite states.
- TestIsInitialStateValid_Invalid - Verifies error when initial state is not found in finite states.
- TestAreTransitionFunctionsValid_Valid - Confirms transition functions are valid with defined states and inputs.
- TestAreTransitionFunctionsValid_InvalidCurrentState - Checks error when a transition function uses an undefined current state.
- TestAreTransitionFunctionsValid_InvalidTransitionState - Verifies error when target state in transition is not in finite states.
- TestAreTransitionFunctionsValid_InvalidInputs - Detects use of input symbols not defined in the automaton's input set.

## 🚀 Getting Started

- All implementation is in pkg\finiteautomation
- example of modulo three implementation using finite automation
  - run the example: go run .\examples\modulothree\main.go
  - modified input in the example for different results
