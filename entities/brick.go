package entities

import (
	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/utils"
)

var totalBricks int = 0

type Brick struct {
	Point   utils.Point
	Visible bool
}

func createBrick(point utils.Point) *Brick {
	brick := &Brick{
		Point:   point,
		Visible: true,
	}
	totalBricks++
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
					Y: float64(5 + y),
				}))
			}
		}
	}

	return bricks
}

func (b *Brick) SetVisibility(vis bool) {
	b.Visible = vis
}

// TODO: (brick collision)

func (b *Brick) Update(ctx *GameContext, dt float64) {
	if !b.Visible {
		return
	}
	const fuzzyThreshold = 1

	brickX := b.Point.X
	brickY := b.Point.Y

	for _, ball := range ctx.Balls {
		ballX := ball.Point.X
		ballY := ball.Point.Y

		if abs(ballX-brickX) <= fuzzyThreshold && abs(ballY-brickY) <= fuzzyThreshold {
			b.SetVisibility(false)
			totalBricks--

			dx := ballX - brickX
			dy := ballY - brickY

			if abs(dx) > abs(dy) {
				if dx > 0 {
					ball.Direction.Left = false
					ball.Direction.Right = true
				} else {
					ball.Direction.Right = false
					ball.Direction.Left = true
				}
			} else {
				if dy > 0 {
					ball.Direction.Up = false
					ball.Direction.Down = true
				} else {
					ball.Direction.Up = true
					ball.Direction.Down = false
				}
			}
		}
	}
}

func (b *Brick) Draw(ctx *GameContext) {
	if b.Visible {
		ctx.Window.SetContent(int(b.Point.X), int(b.Point.Y), 'X')
	}
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
