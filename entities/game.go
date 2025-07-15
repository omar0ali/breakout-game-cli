// Package entities
package entities

import (
	"github.com/omar0ali/breakout-game-cli/core"
)

type GameContext struct {
	Window  *core.Window
	Player  *Player
	Ball    *Ball
	Objects []Entity
}

type Entity interface {
	Update(ctx GameContext, dt float64)
	Draw(ctx GameContext)
}
