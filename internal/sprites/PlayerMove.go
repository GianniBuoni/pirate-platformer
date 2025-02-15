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

func (p *PlayerData) jump() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.direction.Y -= JumpDist
	p.actions["jump"] = true
	p.frameIndex = 0
	go func() {
		wallTimeout := time.NewTimer(400 * time.Millisecond)
		<-wallTimeout.C
		p.actions["wall"] = true
	}()
}

func (p *PlayerData) wallJump(direction float32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.direction.Y -= JumpDist
	p.direction.X = direction
	p.actions["run"] = false
	go func() {
		runTimeout := time.NewTimer(100 * time.Millisecond)
		<-runTimeout.C
		p.actions["run"] = true
	}()
}

func (p *PlayerData) SetGravity(b bool) {
	switch b {
	case true:
		p.gravity = Gravity
	case false:
		p.direction.Y = 0
		p.gravity = Gravity * 0.8
	}
}
