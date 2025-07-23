package entities

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/breakout-game-cli/core"
	"github.com/omar0ali/breakout-game-cli/utils"
)

type Player struct {
	X        float64
	Velocity utils.Velocity

	PlayerSpeed float64
	JumpBy      float64
	PaddleWidth int

	// used when the ball is moving
	Moving   bool
	TargetX  float64
	Traveled float64
}

func CreatePlayer(window *core.Window, config *utils.Config) *Player {
	screenWidth, _ := window.GetScreenSize()
	startPositionOfPaddle := float64(
		(screenWidth / 2) - (int(config.Player.PaddleWdith) / 2),
	) // the middle pos is that start drawing the the bar

	return &Player{
		X:           startPositionOfPaddle,
		Velocity:    utils.Velocity{X: 0, Y: 0},
		PlayerSpeed: config.Player.Speed,
		JumpBy:      config.Player.JumpBy,
		PaddleWidth: int(config.Player.PaddleWdith),
	}
}

func (p *Player) Draw(ctx GameContext) {
	ctx.Debug.AddLine(fmt.Sprintf("Paddle: Moving: %v, PosX: %d", p.Moving, (int(p.X)+p.PaddleWidth)/2))
	_, height := ctx.Window.GetScreenSize()
	for i := 0; i < p.PaddleWidth; i++ {
		ctx.Window.SetContent(int(p.X+float64(i)), height-1, tcell.RuneBlock)
	}
}

func (p *Player) TurnLeft() {
	if p.Moving {
		return
	}
	p.Moving = true
	p.Traveled = 0
	p.Velocity.X = -p.PlayerSpeed
	p.TargetX = p.X - p.JumpBy
}

func (p *Player) TurnRigth() {
	if p.Moving {
		return
	}
	p.Moving = true
	p.Traveled = 0
	p.Velocity.X = p.PlayerSpeed
	p.TargetX = p.X + p.JumpBy
}

func (p *Player) Update(ctx GameContext, dt float64) {
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
	screenWidth, _ := ctx.Window.GetScreenSize()
	if p.X < 0 {
		p.X = 0
		p.Moving = false
	}
	if p.X+float64(p.PaddleWidth) > float64(screenWidth) {
		p.X = float64(screenWidth - p.PaddleWidth)
		p.Moving = false
	}
}
