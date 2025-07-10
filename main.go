package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	window, err := CreateWindow("Breakout Game", 33)
	if err != nil {
		log.Panic(err)
	}
	exit := make(chan int)
	window.InitEventsKeys(
		func(ek *tcell.EventKey) {
			switch ek.Key() {
			// to update an object coordiatnes, not to animate
			case tcell.KeyUp:
				window.SetContent(0, 0, 'O')
			}
		}, func() {
			// animation to draw
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
