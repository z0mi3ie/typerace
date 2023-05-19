package state

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/game"
	"github.com/z0mi3ie/typerace/util"
)

type ScoreState struct {
	isSetup   bool
	completed int
	score     int
	wrong     int
	enabled   bool
	round     *game.Round
}

func (s *ScoreState) setup() {
	s.score = s.round.Score
	s.completed = s.round.Total
	s.wrong = s.round.Incorrect
}

func (s *ScoreState) Update() error {
	if !s.isSetup {
		s.setup()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		stateManager := GetStateManager()
		// TODO: implement some state enums, might be better to have a
		// seperate SceneManager
		// Pop off the extra start state
		stateManager.Pop()
		stateManager.Pop()
		stateManager.Pop()
		return nil
	}

	return nil
}

func (s *ScoreState) Draw(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("Final Score: %d", s.score)
	completedText := fmt.Sprintf("Total words completed: %d", s.completed)
	wrongText := fmt.Sprintf("Incorrect: %d", s.wrong)
	escText := "Press esc to return to main menu..."

	text.Draw(screen,
		scoreText, TextFont,
		(ScreenWidth/2)-util.CenterX(scoreText),
		(ScreenHeight/2)-80,
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
	)

	text.Draw(screen,
		completedText, TextFont,
		(ScreenWidth/2)-util.CenterX(completedText),
		(ScreenHeight/2)-30,
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
	)

	text.Draw(screen,
		wrongText, TextFont,
		(ScreenWidth/2)-util.CenterX(wrongText),
		(ScreenHeight/2)+30,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)

	text.Draw(screen,
		escText, TextFont,
		(ScreenWidth/2)-util.CenterX(escText),
		(ScreenHeight/2)+80,
		color.RGBA{R: 255, G: 50, B: 125, A: 255},
	)
}

func (s *ScoreState) Disable() {
	s.enabled = false
}

func (s *ScoreState) Enable() {
	s.enabled = true
}

func (s *ScoreState) Enabled() bool {
	return s.enabled
}
