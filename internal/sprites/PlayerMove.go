package sprites

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (p *Player) move(dt float32) {
	p.MoveX(p.hitbox, dt)
	p.collison("x")
	p.MoveY(p.hitbox, dt)
	p.collison("y")
	p.patformCollision()
	p.platformMove(dt)
}

func (p *Player) platformMove(dt float32) {
	platform, ok := p.platform.(*MovingSprite)
	if ok {
		platform.MoveX(p.hitbox, dt)
		if platform.direction.Y != 0 {
			p.hitbox.Set(Bottom(platform.hitbox.Top()))
		}
	}
}

func (p *Player) jump() {
	p.platform = nil
	p.direction.Y = -1
	refX := p.hitbox.X
	go func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		timer := time.NewTimer(200 * time.Millisecond)
		<-timer.C
		if p.hitbox.X != refX {
			p.actions[wall] = true
		}
	}()
}

func (p *Player) wallJump(dir float32) {
	if p.actions[wall] {
		p.direction.X = dir
		p.direction.Y -= 1
		p.timeout(run, 100)
	}
}
