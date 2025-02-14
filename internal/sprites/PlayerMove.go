package sprites

import (
	"time"

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

// Player.jump() will make player move vertically up.
// It's also responible for swithing the wall action to true.
// The wall action is deactivated in the Player.checkCollisonSide() method.
func (p *PlayerData) jump() {
	switch p.checkCollisonSide() {
	case floor:
		p.direction.Y -= JumpDist
		go func() {
			wallTimeout := time.NewTimer(400 * time.Millisecond)
			<-wallTimeout.C
			p.actions["wall"] = true
		}()
	}
}
