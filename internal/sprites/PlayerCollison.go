package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionSide uint

const (
	floor CollisionSide = iota
	left
	right
	air
)

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

func (p *PlayerData) platformCollision() {
	for _, s := range *p.platformSprites {
		if rl.CheckCollisionRecs(p.hitbox.Rect(), s.HitBox().Rect()) &&
			p.actions[canPlatform] {
			if p.hitbox.Bottom() >= s.HitBox().Top() &&
				p.oldRect.Bottom() <= s.OldRect().Top() {
				p.hitbox.Set(rects.Bottom(s.HitBox().Top()))
				p.direction.Y = 0
			}
		}
	}
}

func (p *PlayerData) checkCollisonSide() CollisionSide {
	floorRect := rl.NewRectangle(
		p.hitbox.Left()+2, p.hitbox.Bottom(), p.hitbox.Rect().Width-4, 2,
	)
	leftRect := rl.NewRectangle(
		p.hitbox.Left()-2, p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	rightRect := rl.NewRectangle(
		p.hitbox.Right(), p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	for _, s := range *p.collisionSprites {
		if rl.CheckCollisionRecs(floorRect, s.HitBox().Rect()) {
			p.actions[wall] = false
			return floor
		}
		if rl.CheckCollisionRecs(leftRect, s.HitBox().Rect()) &&
			p.actions[wall] {
			p.SetGravity(false)
			return left
		}
		if rl.CheckCollisionRecs(rightRect, s.HitBox().Rect()) &&
			p.actions[wall] {
			p.SetGravity(false)
			return right
		}
	}
	for _, plat := range *p.platformSprites {
		if rl.CheckCollisionRecs(floorRect, plat.HitBox().Rect()) &&
			p.actions[canPlatform] {
			p.actions[wall] = false
			p.actions[platform] = true
			return floor
		}
	}
	p.SetGravity(true)
	return air
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
