package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) move() {
	p.direction.Y += p.gravity
	dt := rl.GetFrameTime()
	// horizontan check
	p.hitbox.X += p.direction.X * dt
	p.collision("x")

	// vertical chek
	p.hitbox.Y += p.direction.Y * dt
	p.collision("y")
}

func (p *PlayerData) collision(axis string) {
	for _, sprite := range *p.collisionSprites {
		if rl.CheckCollisionRecs(sprite.Rect(), p.hitbox) {
			p.gravity = 0
			switch axis {
			case "x":
				if p.Left() <= sprite.Right() {
					p.SetLeft(sprite.Right())
				}
				if p.Right() >= sprite.Left() {
					p.SetRight(sprite.Left())
				}
			case "y":
			}
			if p.Top() <= sprite.Bottom() {
				p.SetTop(sprite.Bottom())
			}
			if p.Bottom() >= sprite.Top() {
				p.SetBottom(sprite.Top())
			}
		}
	}
}
