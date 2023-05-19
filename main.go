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

func startGame() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("typerace")

	stateManager := state.GetStateManager()

	titleState := &state.TitleState{}
	titleState.Load()
	stateManager.Push(titleState)

	err := ebiten.RunGame(stateManager)
	if err != nil {
		log.Fatal(err)
	}
}

func scratch() {
}

func main() {
	//scratch()
	startGame()
}
