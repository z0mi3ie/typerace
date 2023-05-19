package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type StateManager struct {
	states []State
	ii     int
}

var stateMgr *StateManager

func GetStateManager() *StateManager {
	if stateMgr == nil {
		stateMgr = &StateManager{
			ii: -1,
		}
		return stateMgr
	}
	return stateMgr
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) Update() error {
	sm.states[sm.ii].Update()
	return nil
}

func (sm *StateManager) Draw(screen *ebiten.Image) {
	sm.states[sm.ii].Draw(screen)
}

func (sm *StateManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (sm *StateManager) Push(g State) {
	sm.states = append(sm.states, g)
	sm.ii++
}

func (sm *StateManager) Pop() {
	if sm.ii < 1 {
		return
	}
	sm.states = sm.states[:len(sm.states)-1]
	sm.ii--
}
