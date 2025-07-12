package main

import "math"

type Ball struct {
	Point     Point
	Direction Direction
	Velocity  Velocity
}

const BallSpeed float64 = 15

func CreateBall() *Ball {
	width, height := window.GetScreenSize()
	ball := &Ball{
		// placing the ball (middle of the screen) start point
		Point: Point{
			X: float64(width / 2),
			Y: float64(height / 2),
		},
		Direction: Direction{
			Up:    false,
			Down:  true,
			Left:  true,
			Right: false,
		},
		Velocity: Velocity{
			X: 0,
			Y: 0,
		},
	}
	return ball
}

func (b *Ball) Update(dt float64) int {
	width, height := window.GetScreenSize()
	// Move using direction and speed
	b.Velocity.SetFromDirection(BallSpeed, b.Direction.Up, b.Direction.Down, b.Direction.Left, b.Direction.Right)

	b.Point.X += b.Velocity.X * dt
	b.Point.Y += b.Velocity.Y * dt

	// Bounce logic
	if b.Point.Y >= float64(height) {
		playerStartX := float64(player.X - 1)
		playerEndX := player.X + float64(player.PaddleWdith)

		if b.Point.X < playerStartX || b.Point.X > playerEndX {
			return 0 // missed paddle
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
	return 1
}

func (b *Ball) Draw() {
	window.SetContent(int(math.Round(b.Point.X)), int(math.Round(b.Point.Y)), '0')
}
