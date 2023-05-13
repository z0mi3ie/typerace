package state

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/dictionary"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
	"golang.org/x/image/font/basicfont"
)

var (
	TextFont = basicfont.Face7x13
)

type RaceState struct {
	message       string
	inputKeys     []ebiten.Key
	inputCenterX  int
	targetCenterX int
	target        string
	score         int
	dictionary    *dictionary.Dictionary
	soundManager  *sound.SoundManager
	enabled       bool
}

// Load assets from disk and initialize them if needed
func (s *RaceState) Load() {
	s.dictionary = dictionary.New()
	s.soundManager = sound.GetSoundManager()()
}

func (s *RaceState) Update() error {
	if s.target == "" {
		s.target = s.dictionary.Random()
	}

	// Capture the keys and append them to the current keys list
	var ps []ebiten.Key
	pressedKeys := inpututil.AppendJustPressedKeys(ps)
	if input.IsLetterKey(pressedKeys) {
		s.inputKeys = append(s.inputKeys, pressedKeys...)
		s.soundManager.Play("text-input")
	}
	if input.IsBackspaceKey(pressedKeys) {
		if len(s.inputKeys) > 0 {
			s.inputKeys = s.inputKeys[:len(s.inputKeys)-1]
			s.soundManager.Play("text-delete")
		}
	}
	if input.IsEnterKey(pressedKeys) {
		s.inputKeys = []ebiten.Key{}
		// if the current word matches then increase score
		if s.message == s.target {
			s.soundManager.Play("good")
			s.score += 1
			s.inputKeys = []ebiten.Key{}
			s.message = ""
			s.target = ""
		} else {
			s.soundManager.Play("error")
		}
	}

	// Convert the current keys to a displayable string
	var converted string
	for _, k := range s.inputKeys {
		converted = converted + k.String()
	}
	s.message = converted

	// Update the center point of the string to render from
	s.inputCenterX = CenterX(s.message)
	s.targetCenterX = CenterX(s.target)

	// Quit the game on ESC key release
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil

}

func (s *RaceState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"score: %d\ndictionary: %d", s.score, s.dictionary.Length(),
	))
	text.Draw(screen, s.target, TextFont, (ScreenWidth/2)-s.targetCenterX, ScreenHeight/2-20, color.White)
	text.Draw(screen, s.message, TextFont, (ScreenWidth/2)-s.inputCenterX, ScreenHeight/2, color.White)
}

func (s *RaceState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (s *RaceState) Disable() {
	s.enabled = false
}

func (s *RaceState) Enable() {
	s.enabled = true
}

func (s *RaceState) Enabled() bool {
	return s.enabled
}

func CenterX(t string) int {
	rect := text.BoundString(TextFont, t)
	return (rect.Min.X + rect.Max.X) / 2
}
