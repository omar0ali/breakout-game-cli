package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/entities"
	"github.com/omar0ali/breakout-game-cli/utils"
)

func main() {
	window, err := core.CreateWindow("Breakout Game", 33) // frame rate can be changed from here
	if err != nil {
		log.Panic(err)
	}
	exit := make(chan int)
	// gmae config

	cfg, err := utils.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Objects
	player := entities.CreatePlayer(window, cfg)
	ball := entities.CreateBall(window, cfg)
	// TODO: needs refactoring: Creating a list of bricks on the screen. Clean up the code a bit
	var bricks []entities.Brick
	width, _ := window.GetScreenSize()
	half := width / 2
	for i := half / 2; i < half+half/2; i++ {
		bricks = append(bricks, entities.Brick{
			Point: utils.Point{
				X: float64(i),
				Y: float64(4),
			},
			Visible: true,
		})
	}
	ctx := entities.GameContext{
		Window: window,
		Player: player,
		Ball:   ball,
		Bricks: bricks,
		Objects: []entities.Entity{
			player,
			ball,
		},
	}

	// TODO: here we ensure that all the bricks are added to the entity which ensures:
	// 1. Draw() and 2. Update() are called
	for i := 0; i < len(bricks); i++ {
		ctx.Objects = append(ctx.Objects, &bricks[i])
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
			for _, obj := range ctx.Objects {
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
