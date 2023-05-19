package state

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
	"github.com/z0mi3ie/typerace/util"
)

const title = "typeracer"
const author = "github.com/z0mi3ie"
const continueText = "press enter key..."

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
	if input.IsEnterKey(pressedKeys) {
		startState := StartState{}
		startState.Enable()
		stateManager := GetStateManager()
		stateManager.Push(&startState)
	}

	// Quit the game on ESC key release
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	text.Draw(screen,
		title, TextFont,
		(ScreenWidth/2)-util.CenterX(title),
		ScreenHeight/2-40,
		color.White,
	)

	text.Draw(screen,
		author, TextFont,
		(ScreenWidth/2)-util.CenterX(author),
		ScreenHeight/2+10,
		color.White,
	)

	text.Draw(screen,
		continueText, TextFont,
		(ScreenWidth/2)-util.CenterX(continueText),
		ScreenHeight/2+80,
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
