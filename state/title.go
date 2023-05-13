package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
)

type TitleState struct {
	title        string
	author       string
	soundManager *sound.SoundManager
	enabled      bool
}

func (s *TitleState) Load() {
	//s.soundManager = sound.New()
	s.soundManager = sound.GetSoundManager()()
}

func (s *TitleState) Update() error {
	// if enter key is pressed
	//      deactivate this state update
	//      deactivate this state render
	//      get state manaager
	//      create race state
	//      load race state assets (not great here)
	//      push race state on statemanager
	var ps []ebiten.Key
	pressedKeys := inpututil.AppendJustPressedKeys(ps)
	if input.IsEnterKey(pressedKeys) {
		fmt.Println("enter pressed")
		raceState := RaceState{}
		raceState.Load()
		raceState.Enable()
		stateManager := GetStateManager()
		stateManager.Push(&raceState)
	}

	// if esc key is pressed
	//      exit game

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
}

func (s *TitleState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func (s *TitleState) Disable() {
	s.enabled = false
}

func (s *TitleState) Enable() {
	s.enabled = true
}

func (s *TitleState) Enabled() bool {
	return s.enabled
}
