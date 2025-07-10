package main

import "github.com/gdamore/tcell/v2"

type Player struct {
	X      int
	width  int
	window *Window
}

func CreatePlayer(width int, window *Window) *Player {
	screenWidth, _ := window.GetScreenSize()
	middlePos := (screenWidth / 2) - (width / 2) // the middle pos is that start drawing the the bar

	return &Player{
		X:      middlePos,
		width:  width,
		window: window,
	}
}

func (p *Player) Update() {
	_, height := p.window.GetScreenSize()
	for i := 0; i < p.width; i++ {
		p.window.SetContent(p.X+i, height-1, tcell.RuneBlock)
	}
}

// this function is used to update the coordinates of the player, so they can move left or right

func (p *Player) UpdateCoords(x int) {
	p.X = p.X + x
}
