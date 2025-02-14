package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) move() {
	dt := rl.GetFrameTime()
	// horizontan check
	p.hitbox.Set(rects.Left(
		p.hitbox.Left() + p.direction.X*p.speed*dt,
	))
	p.collision("x")

	// vertical check
	p.direction.Y += p.gravity
	p.hitbox.Set(rects.Top(
		p.hitbox.Top() + p.direction.Y*dt,
	))
	p.collision("y")
}

func (p *PlayerData) jump() {
	switch p.checkCollisonSide() {
	case floor:
		p.direction.Y -= JumpDist
	}
}
