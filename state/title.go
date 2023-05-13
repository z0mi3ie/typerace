package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
)

type TitleState struct {
	title        string
	author       string
	enabled      bool
	soundManager *sound.SoundManager
}

func (s *TitleState) Load() {
	s.soundManager = sound.GetSoundManager()
}

func (s *TitleState) Update() error {
	var ps []ebiten.Key
	pressedKeys := inpututil.AppendJustPressedKeys(ps)
	fmt.Println("title-state")
	if input.IsEnterKey(pressedKeys) {
		fmt.Println("enter pressed")
		raceState := RaceState{}
		raceState.Load()
		raceState.Enable()
		stateManager := GetStateManager()
		stateManager.Push(&raceState)
	}

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "title state")
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
