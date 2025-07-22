package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/entities"
	"github.com/omar0ali/breakout-game-cli/utils"
)

func main() {
	// gmae config
	cfg, err := utils.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	window, err := core.CreateWindow("Breakout Game", cfg) // frame rate can be changed from here
	if err != nil {
		log.Panic(err)
	}
	exit := make(chan int)

	// Objects
	player := entities.CreatePlayer(window, cfg)
	ball := entities.CreateBall(window, cfg)
	bricks := entities.CreateBricks(window, cfg)

	ctx := entities.GameContext{
		Window: window,
		Player: player,
		Ball:   ball,
	}

	// add player and ball into the screen (Objects)
	ctx.AddEntities(ball, player)

	// add bricks into the game screen (Objects)
	for i := range len(bricks) {
		ctx.AddEntities(&bricks[i])
	}

	window.InitEventsKeys(
		func(ek *tcell.EventKey, delta float64) {
			switch ek.Key() {
			// to update an object coordiatnes, not to animate
			case tcell.KeyLeft:
				player.TurnLeft()
			case tcell.KeyRight:
				player.TurnRigth()
			}
		}, func(delta float64) {
			// animation to draw
			for _, obj := range ctx.GetObjects() {
				obj.Draw(ctx)
				obj.Update(ctx, delta)
			}
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
