package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
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
	for _, s := range p.groups["collision"] {
		if rl.CheckCollisionRecs(s.HitBox().Rect(), p.hitbox.Rect()) {
			switch axis {
			case "x":
				if p.hitbox.Left() <= s.HitBox().Right() &&
					p.oldRect.Left() >= s.HitBox().Right() {
					p.hitbox.Set(Left(s.HitBox().Right()))
				}
				if p.hitbox.Right() >= s.HitBox().Left() &&
					p.oldRect.Right() <= s.HitBox().Left() {
					p.hitbox.Set(Right(s.HitBox().Left()))
				}
			case "y":
				if p.hitbox.Top() <= s.HitBox().Bottom() &&
					p.oldRect.Top() >= s.HitBox().Bottom() {
					p.hitbox.Set(Top(s.HitBox().Bottom()))
				}
				if p.hitbox.Bottom() >= s.HitBox().Top() &&
					p.oldRect.Bottom() <= s.HitBox().Top() {
					p.hitbox.Set(Bottom(s.HitBox().Top()))
				}
				p.direction.Y = 0
			}
		}
	}
}

func (p *PlayerData) platformCollision() {
	if p.actions[canPlatform] {
		for _, pl := range p.groups["platform"] {
			if rl.CheckCollisionRecs(p.hitbox.Rect(), pl.HitBox().Rect()) {
				if p.hitbox.Bottom() >= pl.HitBox().Top() &&
					p.oldRect.Bottom() <= pl.OldRect().Top() {
					p.hitbox.Set(Bottom(pl.HitBox().Top()))
					p.platform = pl
					p.direction.Y = 0
				}
			}
		}
	}
}

func (p *PlayerData) checkCollisonSide() CollisionSide {
	floorRect := rl.NewRectangle(
		p.hitbox.Left()+2, p.hitbox.Bottom(), p.hitbox.Rect().Width-4, 4,
	)
	leftRect := rl.NewRectangle(
		p.hitbox.Left()-2, p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	rightRect := rl.NewRectangle(
		p.hitbox.Right(), p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	for _, s := range p.groups["collision"] {
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
	if p.platform != nil && p.actions[canPlatform] {
		if rl.CheckCollisionRecs(floorRect, p.platform.HitBox().Rect()) {
			p.actions[wall] = false
			return floor
		} else {
			p.platform = nil
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
