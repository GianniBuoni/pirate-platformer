package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) move() {
	dt := rl.GetFrameTime()
	// horizontan check
	p.hitbox.X += p.direction.X * PlayerSpeed * dt
	p.collision("x")

	// vertical check
	p.direction.Y += p.gravity
	p.hitbox.Y += p.direction.Y * dt
	p.collision("y")
}

func (p *PlayerData) collision(axis string) {
	for _, sprite := range *p.collisionSprites {
		if rl.CheckCollisionRecs(sprite.Rect(), p.hitbox) {
			switch axis {
			case "x":
				fmt.Println(p.direction.X)
				if p.Left() <= sprite.Right() &&
					p.direction.X < 0 {
					p.SetLeft(sprite.Right())
				}
				if p.Right() >= sprite.Left() &&
					p.direction.X > 0 {
					p.SetRight(sprite.Left())
				}
			case "y":
				if p.Top() <= sprite.Bottom() &&
					p.direction.Y < 0 {
					p.SetTop(sprite.Bottom())
				}
				if p.Bottom() >= sprite.Top() &&
					p.direction.Y > 0 {
					p.SetBottom(sprite.Top())
				}
				p.direction.Y = 0
			}
		}
	}
}
