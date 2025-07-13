package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var (
	player *Player
	ball   *Ball
	window *Window
	err    error
)

func main() {
	window, err = CreateWindow("Breakout Game", 33) // frame rate can be changed from here
	if err != nil {
		log.Panic(err)
	}
	exit := make(chan int)

	// Objects
	player = CreatePlayer(10, 50, 5)
	ball = CreateBall(20)

	window.InitEventsKeys(
		func(ek *tcell.EventKey, delta float64) {
			switch ek.Key() {
			// to update an object coordiatnes, not to animate
			case tcell.KeyLeft:
				player.StartMove(-1) // left
			case tcell.KeyRight:
				player.StartMove(1) // right
			}
		}, func(delta float64) {
			// animation to draw
			player.Draw()
			player.Update(delta)

			ball.Draw()
			ball.Update(delta)
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
