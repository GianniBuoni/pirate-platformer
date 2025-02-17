package sprites

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) move() {
	dt := rl.GetFrameTime()
	// horizontan check
	p.hitbox.Set(Left(
		p.hitbox.Left() + p.direction.X*p.speed*dt,
	))
	p.collision("x")

	// vertical check
	p.direction.Y += p.gravity
	p.hitbox.Set(Top(
		p.hitbox.Top() + p.direction.Y*dt,
	))
	p.collision("y")
	p.platformCollision()
}

func (p *PlayerData) attack() {
	if p.actions[canAttack] {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.frameIndex = 0
		p.actions[attack] = true // end of animation toggles this to false
		go p.timeout(canAttack, 400)
	}
}

func (p *PlayerData) jump() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.direction.Y -= JumpDist
	p.actions[jump] = true
	p.frameIndex = 0
	go func() {
		wallTimeout := time.NewTimer(400 * time.Millisecond)
		<-wallTimeout.C
		p.actions[wall] = true
	}()
}

func (p *PlayerData) wallJump(direction float32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.direction.Y -= JumpDist
	p.direction.X = direction
	go p.timeout(run, 100)
}

func (p *PlayerData) phaseThrough() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.actions[platform] = false
	go p.timeout(canPlatform, 200)
}

func (p *PlayerData) timeout(state PlayerState, ms time.Duration) {
	p.actions[state] = !p.actions[state]
	timer := time.NewTimer(ms * time.Millisecond)
	<-timer.C
	p.actions[state] = !p.actions[state]
}
