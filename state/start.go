package state

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/util"
	"image/color"
)

type StartState struct {
	isSetup      bool
	count        *util.Integer
	enabled      bool
	doneCounting chan bool
}

func (s *StartState) setup() {
	fmt.Println("> start state setup")
	s.count = &util.Integer{
		Int: 5,
	}
	s.doneCounting = util.CountDown(s.count, func(n int) {
		fmt.Println("time remaining: ", n)
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
	case _ = <-s.doneCounting:
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
	// Sense we are using a pointer to save the integer
	// being operated on by a ticker, make sure its been
	// initialized before the engine tries to render
	if s.count != nil {
		text.Draw(screen,
			fmt.Sprintf("Get ready... %d", s.count.Int), TextFont,
			(ScreenWidth/2)-util.CenterX(title),
			ScreenHeight/2-40,
			color.White,
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
