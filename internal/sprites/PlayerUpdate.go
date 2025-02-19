package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

func (p *PlayerData) Update() {
	p.oldRect.Copy(p.hitbox)
	p.input()
	p.move()
	p.updateCRects()
	p.rect.Set(
		Top(p.hitbox.Top()-p.hitboxOffset.Y),
		Left(p.hitbox.Left()-p.hitboxOffset.X),
	)
}

func (p *PlayerData) updateCRects() {
	p.cRects[floor].Set(
		Top(p.HitBox().Bottom()), Left(p.HitBox().Left()),
	)
	p.cRects[right].Set(
		Left(p.HitBox().Right()), Top(p.HitBox().Top()+2),
	)
	p.cRects[left].Set(
		Right(p.HitBox().Left()), Top(p.HitBox().Top()+2),
	)
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
