package entities

import (
	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/utils"
)

type Brick struct {
	Point   utils.Point
	Visible bool
}

func createBrick(point utils.Point) *Brick {
	brick := &Brick{
		Point:   point,
		Visible: true,
	}
	return brick
}

func CreateBricks(window *core.Window, cfg *utils.Config) []Brick {
	var bricks []Brick
	width, _ := window.GetScreenSize()
	half := width / 2
	for i := half / 2; i < half+half/2; i++ {
		for y := range cfg.Brick.Level * 2 {
			if y%2 == 0 {
				bricks = append(bricks, *createBrick(utils.Point{
					X: float64(i),
					// start from y = 4 | giving a little gap at the top
					Y: float64(4 + y),
				}))
			}
		}
	}

	return bricks
}

func (b *Brick) SetVisibility(vis bool) {
	b.Visible = vis
}

func (b *Brick) Update(ctx GameContext, dt float64) {
	if int(ctx.Ball.Point.X) == int(b.Point.X) &&
		int(ctx.Ball.Point.Y) == int(b.Point.Y) {
		if b.Visible {
			b.SetVisibility(false)
			if ctx.Ball.Direction.Down {
				ctx.Ball.Direction.Down = false
				ctx.Ball.Direction.Up = true
			} else if ctx.Ball.Direction.Up {
				ctx.Ball.Direction.Down = true
				ctx.Ball.Direction.Up = false
			}
		}
	}
}

func (b *Brick) Draw(ctx GameContext) {
	if b.Visible {
		ctx.Window.SetContent(int(b.Point.X), int(b.Point.Y), 'X')
	}
}
