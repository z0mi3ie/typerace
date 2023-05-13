package state

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type StateManager struct {
	states []State
}

var stateMgr *StateManager

/*
func GetStateManager() func() *StateManager {
	var sm *StateManager
	return func() *StateManager {
		if sm == nil {
			sm = &StateManager{}
			return sm
		}
		return sm
	}
}
*/

func GetStateManager() *StateManager {
	if stateMgr == nil {
		stateMgr = &StateManager{}
		return stateMgr
	}
	return stateMgr
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) Update() error {
	for _, s := range sm.states {
		if s.Enabled() {
			s.Update()
		}
	}
	return nil
}

func (sm *StateManager) Draw(screen *ebiten.Image) {
	for _, s := range sm.states {
		if s.Enabled() {
			s.Draw(screen)
		}
	}
}

func (sm *StateManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (sm *StateManager) Push(g State) {
	// current head disable
	// enable new state
	// push passed state
	if len(sm.states)-1 > 0 {
		sm.states[len(sm.states)-1].Disable()
	}
	sm.states = append(sm.states, g)
	sm.states[len(sm.states)-1].Enable()
	fmt.Println(len(sm.states))
}

func (sm *StateManager) Peak() State {
	return sm.states[len(sm.states)-1]
}

func (sm *StateManager) Pop() State {
	if len(sm.states) < 1 {
		return nil
	}
	// disable current head
	// pop
	// enable new head
	if len(sm.states)-1 > 0 {
		sm.states[len(sm.states)-1].Disable()
	}
	t := sm.states[len(sm.states)-1]
	sm.states = sm.states[:len(sm.states)-1]
	sm.states[len(sm.states)-1].Enable()
	return t
}
