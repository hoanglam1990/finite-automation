package main

import (
	"finite-automation/pkg/finiteautomation/models"
	"fmt"
	"log"
)

func main() {
	// Create a finite set of states
	state1, state2, state3 := models.State{}, models.State{}, models.State{}
	state1.Initialize("0", map[string]*models.State{}) // --> state 1 with output "0"
	state2.Initialize("1", map[string]*models.State{})
	state3.Initialize("2", map[string]*models.State{})

	// Create a set of transition functions
	tf1 := models.TransitionFunction{}
	tf2 := models.TransitionFunction{}
	tf3 := models.TransitionFunction{}
	tf4 := models.TransitionFunction{}
	tf5 := models.TransitionFunction{}
	tf6 := models.TransitionFunction{}

	tf1.Initialize(&state1, "0", &state1) // --> δ(state1,0) = state1
	tf2.Initialize(&state1, "1", &state2)
	tf3.Initialize(&state2, "0", &state3)
	tf4.Initialize(&state2, "1", &state1)
	tf5.Initialize(&state3, "0", &state2)
	tf6.Initialize(&state3, "1", &state3)

	// Create a FiniteAutomation and assgin attributes
	fa := models.FiniteAutomation{}
	err := fa.InitializeFiniteAutomation(
		// Finite set of states
		map[*models.State]*models.State{
			&state1: &state1,
			&state2: &state2,
			&state3: &state3,
		},
		// Finite inputs
		map[string]bool{"0": true, "1": true},
		// Initial state
		&state1,
		// Accepting States
		[]*models.State{&state1, &state2, &state3},
		// Transition functions
		[]models.TransitionFunction{tf1, tf2, tf3, tf4, tf5, tf6},
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	inputs := []string{"1001", "10110", "101100"}
	for _, input := range inputs {
		fmt.Print("Input: ", input, " - ")
		result, err := fa.Compute(input)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("Final state output: ", *result)
	}
}
