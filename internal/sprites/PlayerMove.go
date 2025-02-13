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
		p.hitbox.Left() + p.direction.X*PlayerSpeed*dt,
	))
	p.collision("x")

	// vertical check
	p.direction.Y += p.gravity
	p.hitbox.Set(rects.Top(
		p.hitbox.Top() + p.direction.Y*dt,
	))
	p.collision("y")
}

func (p *PlayerData) collision(axis string) {
	for _, s := range *p.collisionSprites {
		if rl.CheckCollisionRecs(s.HitBox().Rect(), p.hitbox.Rect()) {
			switch axis {
			case "x":
				if p.hitbox.Left() <= s.HitBox().Right() &&
					p.oldRect.Left() >= s.HitBox().Right() {
					p.hitbox.Set(rects.Left(s.HitBox().Right()))
				}
				if p.hitbox.Right() >= s.HitBox().Left() &&
					p.oldRect.Right() <= s.HitBox().Left() {
					p.hitbox.Set(rects.Right(s.HitBox().Left()))
				}
			case "y":
				if p.hitbox.Top() <= s.HitBox().Bottom() &&
					p.oldRect.Top() >= s.HitBox().Bottom() {
					p.hitbox.Set(rects.Top(s.HitBox().Bottom()))
				}
				if p.hitbox.Bottom() >= s.HitBox().Top() &&
					p.oldRect.Bottom() <= s.HitBox().Top() {
					p.hitbox.Set(rects.Bottom(s.HitBox().Top()))
				}
				p.direction.Y = 0
			}
		}
	}
}
