package entities

import (
	"math"

	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/utils"
)

type Ball struct {
	Point     utils.Point
	Direction utils.Direction
	Velocity  utils.Velocity
	BallSpeed float64
}

func (b *Ball) SetBallPosition(point utils.Point) {
	b.Point = point
}

func (b *Ball) ResetBallPosition(ctx GameContext) {
	width, height := ctx.Window.GetScreenSize()
	b.SetBallPosition(utils.Point{
		X: float64(width / 2),
		Y: float64(height / 2),
	})
}

func CreateBall(window *core.Window, config *utils.Config) *Ball {
	width, height := window.GetScreenSize()
	ball := &Ball{
		// placing the ball (middle of the screen) start point
		Point: utils.Point{
			X: float64(width / 2),
			Y: float64(height / 2),
		},
		Direction: utils.Direction{
			Up:    false,
			Down:  true,
			Left:  true,
			Right: false,
		},
		Velocity: utils.Velocity{
			X: 0,
			Y: 0,
		},
		BallSpeed: config.Ball.Speed,
	}
	return ball
}

func (b *Ball) Update(ctx GameContext, dt float64) {
	width, height := ctx.Window.GetScreenSize()
	b.Velocity.SetFromDirection(b.BallSpeed, b.Direction.Up, b.Direction.Down, b.Direction.Left, b.Direction.Right)

	b.Point.X += b.Velocity.X * dt
	b.Point.Y += b.Velocity.Y * dt

	// bounce logic
	if b.Point.Y >= float64(height) {
		playerStartX := float64(ctx.Player.X - 1)
		playerEndX := ctx.Player.X + float64(ctx.Player.PaddleWidth)

		// ball fall over the paddle
		if b.Point.X < playerStartX || b.Point.X > playerEndX {
			b.ResetBallPosition(ctx)
		}
		b.Direction.Down = false
		b.Direction.Up = true
	}
	if b.Point.Y <= 0 {
		b.Direction.Up = false
		b.Direction.Down = true
	}
	if b.Point.X <= 0 {
		b.Direction.Left = false
		b.Direction.Right = true
	}
	if b.Point.X >= float64(width-1) {
		b.Direction.Left = true
		b.Direction.Right = false
	}
}

func (b *Ball) Draw(ctx GameContext) {
	ctx.Window.SetContent(int(math.Round(b.Point.X)), int(math.Round(b.Point.Y)), '0')
}
