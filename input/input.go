package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	asciiLetterKeys = []ebiten.Key{
		ebiten.KeyA,
		ebiten.KeyB,
		ebiten.KeyC,
		ebiten.KeyD,
		ebiten.KeyE,
		ebiten.KeyF,
		ebiten.KeyG,
		ebiten.KeyH,
		ebiten.KeyI,
		ebiten.KeyJ,
		ebiten.KeyK,
		ebiten.KeyL,
		ebiten.KeyM,
		ebiten.KeyN,
		ebiten.KeyO,
		ebiten.KeyP,
		ebiten.KeyQ,
		ebiten.KeyR,
		ebiten.KeyS,
		ebiten.KeyT,
		ebiten.KeyU,
		ebiten.KeyV,
		ebiten.KeyW,
		ebiten.KeyX,
		ebiten.KeyY,
		ebiten.KeyZ,
	}
)

func IsLetterKey(k []ebiten.Key) bool {
	if len(k) != 1 {
		return false
	}
	for _, l := range asciiLetterKeys {
		if k[0] == l {
			return true
		}
	}
	return false
}

func IsEnterKey(k []ebiten.Key) bool {
	if len(k) != 1 {
		return false
	}
	return k[0] == ebiten.KeyEnter
}

func IsBackspaceKey(k []ebiten.Key) bool {
	if len(k) != 1 {
		return false
	}
	return k[0] == ebiten.KeyBackspace
}
