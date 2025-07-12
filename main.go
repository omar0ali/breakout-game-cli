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
	player = CreatePlayer(10)
	ball = CreateBall()

	window.InitEventsKeys(
		func(ek *tcell.EventKey) {
			switch ek.Key() {
			// to update an object coordiatnes, not to animate
			case tcell.KeyLeft:
				player.UpdateCoords(-5) // left
			case tcell.KeyRight:
				player.UpdateCoords(5) // right
			}
		}, func(delta float64) {
			// animation to draw
			player.Draw()
			ball.Draw()
			if ball.Update(delta) == 0 { // game over
				window.Close()
				exit <- 0
				return
			}
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
