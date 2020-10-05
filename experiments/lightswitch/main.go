package main

import "fmt"

const (
	Off StateType = "Off"
	On  StateType = "On"

	SwitchOff EventType = "SwitchOff"
	SwitchOn  EventType = "SwitchOn"
)

// OffAction represents the action executed on entering the Off state.
type OffAction struct{}

func (a *OffAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("The light has been switched off")
	return NoOp
}

// OnAction represents the action executed on entering the On state.
type OnAction struct{}

func (a *OnAction) Execute(eventCtx EventContext) EventType {
	fmt.Println("The light has been switched on")
	return NoOp
}

func newLightSwitchFSM() *StateMachine {
	return &StateMachine{
		States: States{
			Default: State{
				Events: Events{
					SwitchOff: Off,
				},
			},
			Off: State{
				Action: &OffAction{},
				Events: Events{
					SwitchOn: On,
				},
			},
			On: State{
				Action: &OnAction{},
				Events: Events{
					SwitchOff: Off,
				},
			},
		},
	}
}
