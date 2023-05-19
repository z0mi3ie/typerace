package util

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"time"
)

var (
	TextFont = basicfont.Face7x13
)

func CenterX(t string) int {
	rect := text.BoundString(TextFont, t)
	return (rect.Min.X + rect.Max.X) / 2
}

type TickAction func(int)

type Integer struct {
	Int int
}

func CountDown(n *Integer, f TickAction) chan bool {
	// tick every one second
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	cd := n
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				n.Int--
				f(n.Int)
				if cd.Int == 0 {
					ticker.Stop()
					done <- true
				}
			}
		}
	}()

	fmt.Println("returning channel")
	return done

	/*
		for {
			if cd == 0 {
				ticker.Stop()
				done <- true
				break
			}
		}
	*/
}
