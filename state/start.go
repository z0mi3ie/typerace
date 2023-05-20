package state

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/sound"
	"github.com/z0mi3ie/typerace/util"
)

type StartState struct {
	isSetup bool
	count   *util.Integer
	enabled bool
	done    chan bool
}

func (s *StartState) setup() {
	fmt.Println("> start state setup")
	sndMgr := sound.GetSoundManager()
	sndMgr.Load()

	s.count = &util.Integer{
		Int: 5,
	}
	s.done = util.CountDown(s.count, func(n int) {
		fmt.Println("time remaining: ", n)
		if n < 4 {
			sndMgr.Play("good")
		}

	})
	s.isSetup = true
}

func (s *StartState) Update() error {
	if !s.isSetup {
		s.setup()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		stateManager := GetStateManager()
		stateManager.Pop()
		return nil
	}

	select {
	case _ = <-s.done:
		raceState := RaceState{}
		raceState.Load()
		raceState.Enable()
		stateManager := GetStateManager()
		stateManager.Push(&raceState)
	default:
		return nil
	}

	return nil
}

func (s *StartState) Draw(screen *ebiten.Image) {
	// We are using a pointer to save the integer
	// being operated on by a ticker, make sure its been
	// initialized before the engine tries to render
	if s.count != nil {
		readyText := "Get READY!"
		timeRemaining := fmt.Sprintf("%d", s.count.Int)
		text.Draw(screen,
			readyText, TextFont,
			(ScreenWidth/2)-util.CenterX(readyText),
			ScreenHeight/2-20,
			color.RGBA{R: 255, G: 0, B: 0, A: 255},
		)

		text.Draw(screen,
			timeRemaining, TextFont,
			(ScreenWidth/2)-util.CenterX(timeRemaining),
			ScreenHeight/2,
			color.RGBA{R: 255, G: 0, B: 0, A: 255},
		)
	}
}

func (s *StartState) Disable() {
	s.enabled = false
}

func (s *StartState) Enable() {
	s.enabled = true
}

func (s *StartState) Enabled() bool {
	return s.enabled
}
