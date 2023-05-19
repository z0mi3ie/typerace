package state

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/dictionary"
	"github.com/z0mi3ie/typerace/game"
	"github.com/z0mi3ie/typerace/input"
	"github.com/z0mi3ie/typerace/sound"
	"github.com/z0mi3ie/typerace/util"
	"golang.org/x/image/font/basicfont"
)

var (
	TextFont = basicfont.Face7x13
)

type RaceState struct {
	isSetup       bool
	message       string
	inputKeys     []ebiten.Key
	inputCenterX  int
	targetCenterX int
	target        string
	total         int
	score         int
	dictionary    *dictionary.Dictionary
	soundManager  *sound.SoundManager
	enabled       bool
	done          chan bool
	count         *util.Integer
	round         *game.Round
}

func (s *RaceState) setup() {
	fmt.Println("> race state setup")
	s.count = &util.Integer{
		Int: 10,
	}
	s.done = util.CountDown(s.count, func(n int) {
		fmt.Println("time remaining: ", n)
	})

	s.round = game.GetRound()
	s.round.Reset()
	s.isSetup = true
}

// Load assets from disk and initialize them if needed
func (s *RaceState) Load() {
	s.dictionary = dictionary.New()
	s.soundManager = sound.GetSoundManager()
	s.soundManager.Load()
}

func (s *RaceState) Update() error {
	if !s.isSetup {
		s.setup()
	}

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
			s.total, s.score = s.round.Correct(s.message)
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
	s.inputCenterX = util.CenterX(s.message)
	s.targetCenterX = util.CenterX(s.target)

	// Go back to the title screen on escape press
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		stateManager := GetStateManager()
		stateManager.Pop()
		// TODO: implement some state enums, might be better to have a combination here
		// Pop off the extra start state
		stateManager.Pop()
		return nil
	}

	select {
	case _ = <-s.done:
		// TODO: round has ended and end round state should be present here
		// TODO: go back to title for now
		stateManager := GetStateManager()
		stateManager.Pop()
		// TODO: implement some state enums, might be better to have a combination here
		// Pop off the extra start state
		stateManager.Pop()
		return nil
	default:
		return nil
	}
}

func (s *RaceState) Draw(screen *ebiten.Image) {
	score := fmt.Sprintf("Score: %d", s.score)
	text.Draw(screen,
		score, TextFont,
		(ScreenWidth/2)-util.CenterX(score),
		ScreenHeight/2+80,
		color.White,
	)

	// Current word from dictionary
	text.Draw(screen,
		s.target, TextFont,
		(ScreenWidth/2)-s.targetCenterX,
		ScreenHeight/2-20,
		color.White,
	)

	// Current message typed by user
	text.Draw(screen,
		s.message, TextFont,
		(ScreenWidth/2)-s.inputCenterX,
		ScreenHeight/2,
		color.White,
	)

	// Time remaining
	if s.count != nil {
		timeRemaining := fmt.Sprintf("< %d >", s.count.Int)
		text.Draw(screen,
			timeRemaining, TextFont,
			(ScreenWidth/2)-util.CenterX(timeRemaining),
			30,
			color.White,
		)
	}
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
