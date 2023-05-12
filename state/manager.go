package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type StateManager struct {
	states []State
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) Update() error {
	for _, s := range sm.states {
		s.Update()
	}
	return nil
}

func (sm *StateManager) Draw(screen *ebiten.Image) {
	for _, s := range sm.states {
		s.Draw(screen)
	}
}

func (sm *StateManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (sm *StateManager) Push(g State) {
	sm.states = append(sm.states, g)
}

func (sm *StateManager) Peak() State {
	return sm.states[len(sm.states)-1]
}

func (sm *StateManager) Pop() State {
	if len(sm.states) < 1 {
		return nil
	}

	t := sm.states[len(sm.states)-1]
	sm.states = sm.states[:len(sm.states)-1]
	return t
}
