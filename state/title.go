package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type TitleState struct{}

func (s TitleState) Update() error {
	fmt.Println("TITLE STATE")
	return nil
}

func (s TitleState) Draw(screen *ebiten.Image) {
}

func (s TitleState) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
