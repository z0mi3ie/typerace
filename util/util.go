package util

import (
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var (
	TextFont = basicfont.Face7x13
)

func CenterX(t string) int {
	rect := text.BoundString(TextFont, t)
	return (rect.Min.X + rect.Max.X) / 2
}
