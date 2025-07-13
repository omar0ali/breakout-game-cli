package main

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	X        float64
	Velocity Velocity

	PlayerSpeed float64
	JumpBy      float64
	PaddleWidth int

	// used when the ball is moving
	Moving   bool
	TargetX  float64
	Traveled float64
}

func CreatePlayer(paddleWidth int, playerSpeed, jumpBy float64) *Player {
	screenWidth, _ := window.GetScreenSize()
	startPositionOfPaddle := float64((screenWidth / 2) - (paddleWidth / 2)) // the middle pos is that start drawing the the bar

	return &Player{
		X:           startPositionOfPaddle,
		Velocity:    Velocity{X: 0, Y: 0},
		PlayerSpeed: playerSpeed,
		JumpBy:      jumpBy,
		PaddleWidth: paddleWidth,
	}
}

func (p *Player) Draw() {
	_, height := window.GetScreenSize()
	for i := 0; i < p.PaddleWidth; i++ {
		window.SetContent(int(p.X+float64(i)), height-1, tcell.RuneBlock)
	}
}

func (p *Player) StartMove(dir int) {
	if p.Moving {
		return
	}

	p.Moving = true
	p.Traveled = 0

	if dir < 0 {
		p.Velocity.X = -p.PlayerSpeed
		p.TargetX = p.X - p.JumpBy
	} else {
		p.Velocity.X = p.PlayerSpeed
		p.TargetX = p.X + p.JumpBy
	}
}

func (p *Player) Update(dt float64) {
	if !p.Moving {
		return
	}

	step := p.Velocity.X * dt
	p.X += step
	p.Traveled += math.Abs(step)

	if p.Traveled >= p.JumpBy {
		p.X = p.TargetX
		p.Moving = false
		p.Velocity.X = 0
	}
	// boundary check
	screenWidth, _ := window.GetScreenSize()
	if p.X < 0 {
		p.X = 0
		p.Moving = false
	}
	if p.X+float64(p.PaddleWidth) > float64(screenWidth) {
		p.X = float64(screenWidth - p.PaddleWidth)
		p.Moving = false
	}
}
