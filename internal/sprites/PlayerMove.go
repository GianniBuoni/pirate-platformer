package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

func (p *Player) move(dt float32) {
	p.MoveX(p.hitbox, dt)
	p.collison("x")
	p.MoveY(p.hitbox, dt)
	p.collison("y")
}

func (p *Player) jump() {
	p.direction.Y -= JumpDist
}
