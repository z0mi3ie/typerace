package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	Disable()
	Enable()
	Enabled() bool
}
