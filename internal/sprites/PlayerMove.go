package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (p *Player) move(dt float32) {
	p.MoveX(p.hitbox, dt)
	p.collison("x")
	p.MoveY(p.hitbox, dt)
	p.collison("y")
	p.patformCollision()
	p.platformMove(dt)
	p.damageCollision()
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
	p.frameIndex = 0
	p.actions[jump] = true
	p.timeout(wall, 200)
}

func (p *Player) wallJump(dir float32) {
	if p.actions[wall] {
		p.direction.X = dir
		p.direction.Y -= 1
		p.timeout(run, 100)
	}
}

func (p *Player) attack() {
	if p.actions[canAttack] {
		p.frameIndex = 0
		p.actions[attack] = true
		p.timeout(canAttack, 600)
	}
}

func (p *Player) phasePlatform() {
	p.platform = nil
	p.timeout(canPlatform, 200)
}
