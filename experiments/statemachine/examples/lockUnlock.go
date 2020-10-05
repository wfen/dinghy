package main

import (
	"fmt"

	"github.com/wfen/dingy/experiments/statemachine"
)

func main() {
	machine := statemachine.Machine{
		ID:      "lock-unlock",
		Initial: "locked",
		States: statemachine.StateMap{
			"locked": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"COIN": statemachine.MachineTransition{
						To: "unlocked",
					},
					"PUSH": statemachine.MachineTransition{
						To: "locked",
					},
				},
			},
			"unlocked": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"COIN": statemachine.MachineTransition{
						To: "unlocked",
					},
					"PUSH": statemachine.MachineTransition{
						To: "locked",
					},
				},
			},
		},
	}
	fmt.Printf("current: %v \t", machine.Current())
	output := machine.Transition("COIN")
	fmt.Printf("next: %v \n", output)

	fmt.Printf("current: %v \t", machine.Current())
	output = machine.Transition("COIN")
	fmt.Printf("next: %v \n", output)

	fmt.Printf("current: %v \t", machine.Current())
	output = machine.Transition("PUSH")
	fmt.Printf("next: %v \n", output)

	fmt.Printf("current: %v \t", machine.Current())
	output = machine.Transition("COIN")
	fmt.Printf("next: %v \n", output)

}
