package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/z0mi3ie/typerace/state"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("typerace")

	raceState := &state.RaceState{}
	raceState.Load()

	stateManager := state.NewStateManager()
	stateManager.Push(raceState)

	err := ebiten.RunGame(stateManager)
	if err != nil {
		log.Fatal(err)
	}
}
