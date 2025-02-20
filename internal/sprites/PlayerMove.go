package sprites

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) move() {
	dt := rl.GetFrameTime()
	// horizontan check
	p.hitbox.X += p.direction.X * p.speed * dt
	p.collision("x")

	// vertical check
	p.direction.Y += p.gravity * dt
	p.hitbox.Y += p.direction.Y * dt

	p.collision("y")
	p.platformCollision()
	if p.platform != nil {
		if vert := p.platform.Movement(p.hitbox); vert {
			p.hitbox.Set(Bottom(p.platform.HitBox().Top()))
		}
	}
	p.hitCollision()
}

func (p *PlayerData) attack() {
	if p.actions[canAttack] {
		p.frameIndex = 0
		p.mu.RLock()
		defer p.mu.RUnlock()
		p.actions[attack] = true // end of animation toggles this to false
		p.timeout(canAttack, 400)
	}
}

func (p *PlayerData) jump() {
	p.platform = nil
	p.direction.Y -= JumpDist
	p.frameIndex = 0
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.actions[jump] = true
	go func() {
		wallTimeout := time.NewTimer(400 * time.Millisecond)
		<-wallTimeout.C
		p.actions[wall] = true
	}()
}

func (p *PlayerData) wallJump(direction float32) {
	p.direction.Y -= JumpDist
	p.direction.X = direction
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.timeout(run, 100)
}

func (p *PlayerData) phaseThrough() {
	p.platform = nil
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.timeout(canPlatform, 200)
}

func (p *PlayerData) timeout(state PlayerState, ms time.Duration) {
	p.actions[state] = !p.actions[state]
	go func() {
		timer := time.NewTimer(ms * time.Millisecond)
		<-timer.C
		p.actions[state] = !p.actions[state]
	}()
}
