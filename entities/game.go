// Package entities
package entities

import (
	"github.com/omar0ali/breakout-game-cli/core"
)

// GameContext shares objects on the screen for easier access
// Bricks not required here since the ball nor the player needs to do anything to it

type GameContext struct {
	Window  *core.Window
	Player  *Player
	Balls   map[int]*Ball
	objects []Entity
	Debug   *core.Debug
}

type Entity interface {
	Update(ctx *GameContext, dt float64)
	Draw(ctx *GameContext)
}

func (ctx *GameContext) AddEntities(entities ...Entity) {
	ctx.objects = append(ctx.objects, entities...)
}

func (ctx *GameContext) RemoveEntity(en Entity) {
	for i, x := range ctx.objects {
		if x == en {
			// Once found, remove it
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GameContext) GetObjects() []Entity {
	return ctx.objects
}
