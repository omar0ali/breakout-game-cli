package main

type Ball struct {
	X      int
	Y      int
	window *Window
	Up     bool
	Down   bool
	Player *Player
}

func CreateBall(window *Window, player *Player) *Ball {
	width, height := window.GetScreenSize()
	ball := &Ball{
		X:      width / 2,
		Y:      height / 2,
		window: window,
		Up:     false,
		Down:   true,
		Player: player,
	}
	return ball
}

func (b *Ball) Update() int {
	_, height := b.window.GetScreenSize()
	if b.Down {
		b.Y = b.Y + 1
		if b.Y >= height-1 {
			startPos := b.Player.X
			endPos := b.Player.X + b.Player.width - 1

			if b.X < startPos || b.X > endPos {
				// Missed paddle — Game over
				return 0
			}

			b.Down = false
			b.Up = true
		}
	}
	if b.Up {
		b.Y = b.Y - 1
		if b.Y <= 0 {
			b.Down = true
			b.Up = false
		}
	}
	return 1
}

func (b *Ball) Draw() {
	b.window.SetContent(b.X, b.Y, '0')
}
