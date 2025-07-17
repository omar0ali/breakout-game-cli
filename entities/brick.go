package entities

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/breakout-game-cli/utils"
)

type Brick struct {
	Point   utils.Point
	Visible bool
}

func CreateBrick(point utils.Point) *Brick {
	brick := &Brick{
		Point:   point,
		Visible: true,
	}
	return brick
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
		ctx.Window.SetContent(int(b.Point.X), int(b.Point.Y), tcell.RuneCkBoard)
	}
}
