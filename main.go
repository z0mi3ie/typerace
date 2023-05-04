package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/z0mi3ie/typerace/dictionary"
	"golang.org/x/image/font/basicfont"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	SampleRate   = 48000
	SoundsDir    = "assets/sounds"
)

var (
	TextFont = basicfont.Face7x13
)

type Game struct {
	message       string
	inputKeys     []ebiten.Key
	inputCenterX  int
	targetCenterX int
	target        string
	score         int
	dictionary    *dictionary.Dictionary
	audioContext  *audio.Context
	audioPlayer   *audio.Player
}

func CenterX(t string) int {
	rect := text.BoundString(TextFont, t)
	return (rect.Min.X + rect.Max.X) / 2
}

func (g *Game) Load() {
	g.dictionary = dictionary.New()
	g.audioContext = audio.NewContext(48000)

	textSound := MustLoadSound("text_1_16.wav")
	p, err := g.audioContext.NewPlayer(textSound)
	if err != nil {
		log.Panicf(err.Error())
	}
	g.audioPlayer = p
}

func (g *Game) Update() error {
	if g.target == "" {
		g.target = g.dictionary.Random()
	}

	// Capture the keys and append them to the current keys list
	var ps []ebiten.Key
	pressedKeys := inpututil.AppendJustPressedKeys(ps)
	if IsLetterKey(pressedKeys) {
		g.inputKeys = append(g.inputKeys, pressedKeys...)
		g.audioPlayer.Rewind()
		g.audioPlayer.Play()
	}
	if IsBackspaceKey(pressedKeys) {
		if len(g.inputKeys) > 0 {
			g.inputKeys = g.inputKeys[:len(g.inputKeys)-1]
			g.audioPlayer.Rewind()
			g.audioPlayer.Play()
		}
	}
	if IsEnterKey(pressedKeys) {
		g.inputKeys = []ebiten.Key{}
		g.audioPlayer.Rewind()
		g.audioPlayer.Play()
		// if the current word matches then increase score
		if g.message == g.target {
			g.score += 1
			g.inputKeys = []ebiten.Key{}
			g.message = ""
			g.target = ""
		}
	}

	// Convert the current keys to a displayable string
	var converted string
	for _, k := range g.inputKeys {
		converted = converted + k.String()
	}
	g.message = converted

	// Update the center point of the string to render from
	g.inputCenterX = CenterX(g.message)
	g.targetCenterX = CenterX(g.target)

	// Quit the game on ESC key release
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"score: %d\ndictionary: %d", g.score, g.dictionary.Length(),
	))
	text.Draw(screen, g.target, TextFont, (ScreenWidth/2)-g.targetCenterX, ScreenHeight/2-20, color.White)
	text.Draw(screen, g.message, TextFont, (ScreenWidth/2)-g.inputCenterX, ScreenHeight/2, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func MustLoadSound(n string) *wav.Stream {
	tf := fmt.Sprintf("./%s/%s", SoundsDir, n)
	f, err := os.Open(tf)
	if err != nil {
		log.Panicf("Could not open file %s", tf)
	}
	d, err := wav.DecodeWithoutResampling(f)
	if err != nil {
		log.Panicf("Could not decode file %s: %s", tf, err.Error())
	}
	return d
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("typerace")

	game := &Game{}
	game.Load()

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
