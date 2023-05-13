package state

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
	"github.com/z0mi3ie/typerace/util"
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

	// Quit the game on ESC key release
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "title state")

	text.Draw(screen,
		"typracer", TextFont,
		(ScreenWidth/2)-util.CenterX("typeracer"),
		ScreenHeight/2-20,
		color.White,
	)

	text.Draw(screen,
		"press enter key...", TextFont,
		(ScreenWidth/2)-util.CenterX("press enter key..."),
		ScreenHeight/2+20,
		color.White,
	)
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
