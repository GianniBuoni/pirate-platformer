package sprites

import (
	"fmt"

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
	for _, sprite := range *p.collisionSprites {
		if rl.CheckCollisionRecs(sprite.HitBox().Rect(), p.HitBox().Rect()) {
			switch axis {
			case "x":
				fmt.Println(p.direction.X)
				if p.hitbox.Left() <= sprite.HitBox().Right() &&
					p.direction.X < 0 {
					p.hitbox.Set(rects.Left(sprite.HitBox().Right()))
				}
				if p.hitbox.Right() >= sprite.HitBox().Left() &&
					p.direction.X > 0 {
					p.hitbox.Set(rects.Right(sprite.HitBox().Left()))
				}
			case "y":
				if p.hitbox.Top() <= sprite.HitBox().Bottom() &&
					p.direction.Y < 0 {
					p.hitbox.Set(rects.Top(sprite.HitBox().Bottom()))
				}
				if p.hitbox.Bottom() >= sprite.HitBox().Top() &&
					p.direction.Y > 0 {
					p.hitbox.Set(rects.Bottom(sprite.HitBox().Top()))
				}
				p.direction.Y = 0
			}
		}
	}
}
