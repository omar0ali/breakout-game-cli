package main

type Ball struct {
	window    *Window
	Player    *Player
	point     Point
	direction Direction
	velocity  Velocity
}

func CreateBall(window *Window, player *Player) *Ball {
	width, height := window.GetScreenSize()
	ball := &Ball{
		point: Point{
			X: float64(width / 2),
			Y: float64(height / 2),
		},
		window: window,
		direction: Direction{
			Up:    false,
			Down:  true,
			Left:  true,
			Right: false,
		},
		Player: player,
	}
	return ball
}

func (b *Ball) Update() int {
	width, height := b.window.GetScreenSize()
	if b.direction.Down {
		b.point.Y = b.point.Y + 1
		if b.point.Y >= float64(height-1) {
			startPos := b.Player.X
			endPos := b.Player.X + b.Player.width - 1

			if b.point.X < float64(startPos) || b.point.X > float64(endPos) {
				// Missed paddle — Game over
				return 0
			}
			b.direction.Up = true
			b.direction.Down = false
		}
	}
	if b.direction.Up {
		b.point.Y = b.point.Y - 1
		if b.point.Y <= 0 {
			b.direction.Down = true
			b.direction.Up = false
		}
	}
	if b.direction.Left {
		b.point.X = b.point.X - 1
		if b.point.X <= 0 {
			b.direction.Right = true
			b.direction.Left = false
		}
	}
	if b.direction.Right {
		b.point.X = b.point.X + 1
		if b.point.X >= float64(width) {
			b.direction.Right = false
			b.direction.Left = true
		}
	}
	return 1
}

func (b *Ball) Draw() {
	b.window.SetContent(int(b.point.X), int(b.point.Y), '0')
}
