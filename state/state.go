package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	Disable()
	Enable()
	Enabled() bool
}
