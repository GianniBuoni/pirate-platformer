package sprites

import "time"

func (p *Player) move(dt float32) {
	p.MoveX(p.hitbox, dt)
	p.collison("x")
	p.MoveY(p.hitbox, dt)
	p.collison("y")
	p.patformCollision()
}

func (p *Player) jump() {
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
