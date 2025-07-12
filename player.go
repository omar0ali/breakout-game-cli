package main

import "github.com/gdamore/tcell/v2"

type Player struct {
	X           float64
	PaddleWdith int
	Direction   Direction
}

const PlayerSpeed float64 = 10

func CreatePlayer(paddleWdith int) *Player {
	screenWidth, _ := window.GetScreenSize()
	middlePos := float64((screenWidth / 2) - (paddleWdith / 2)) // the middle pos is that start drawing the the bar

	return &Player{
		X:           middlePos,
		PaddleWdith: paddleWdith,
	}
}

func (p *Player) Draw() {
	_, height := window.GetScreenSize()
	for i := 0; i < p.PaddleWdith; i++ {
		window.SetContent(int(p.X+float64(i)), height-1, tcell.RuneBlock)
	}
}

// this function is used to update the coordinates of the player, so they can move left or right

func (p *Player) UpdateCoords(x float64) {
	width, _ := window.GetScreenSize()

	if p.X+float64(p.PaddleWdith) > float64(width) {
		p.X = float64(width - p.PaddleWdith)
		return
	} else if p.X < 0 {
		p.X = 0
		return
	}
	p.X = p.X + x
}
