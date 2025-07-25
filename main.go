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

	if cfg.Core.Mouse {
		window.Screen.EnableMouse()
	}

	// Objects
	balls := make(map[int]*entities.Ball)
	player := entities.CreatePlayer(window, cfg)
	bricks := entities.CreateBricks(window, cfg)
	debug := core.CreateDebug(utils.Point{
		X: 0, Y: 1,
	}, cfg)

	ctx := &entities.GameContext{
		Window: window,
		Player: player,
		Balls:  balls,
		Debug:  &debug,
	}

	// add player and ball into the screen (Objects)
	ctx.AddEntities(player)

	// add bricks into the game screen (Objects)
	for i := range len(bricks) {
		ctx.AddEntities(&bricks[i])
	}

	window.InitEventsKeys(
		func(ek tcell.Event, delta float64) {
			switch ev := ek.(type) {
			// to update an object coordiatnes, not to animate
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyLeft:
					player.TurnLeft()
				case tcell.KeyRight:
					player.TurnRigth()
				}
				if ev.Rune() == ' ' {
					player.ShootBall(ctx, cfg)
				}
			case *tcell.EventMouse:
				x, y := ev.Position()
				player.SetPosition(x, y)
			}
		}, func(delta float64) {
			// animation to draw
			for _, obj := range ctx.GetObjects() {
				obj.Draw(ctx)
				obj.Update(ctx, delta)
			}
			debug.Draw(window)
		}, exit,
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}
