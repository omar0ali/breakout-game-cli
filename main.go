package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var player *Player

func main() {
	window, err := CreateWindow("Breakout Game", 25)
	if err != nil {
		log.Panic(err)
	}
	exit := make(chan int)

	// Objects
	player := CreatePlayer(10, window)

	window.InitEventsKeys(
		func(ek *tcell.EventKey) {
			switch ek.Key() {
			// to update an object coordiatnes, not to animate
			case tcell.KeyUp:
				window.SetContent(0, 0, 'O')
			case tcell.KeyLeft:
				player.UpdateCoords(-2) // left
			case tcell.KeyRight:
				player.UpdateCoords(2) // right
			}
		}, func() {
			// animation to draw
			player.Update()
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
