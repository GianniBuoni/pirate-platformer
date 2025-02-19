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
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, s := range p.groups["collision"] {
		if rl.CheckCollisionRecs(p.cRects[floor].Rect(), s.HitBox().Rect()) {
			p.actions[wall] = false
			return floor
		}
		if rl.CheckCollisionRecs(p.cRects[left].Rect(), s.HitBox().Rect()) &&
			p.actions[wall] {
			p.SetGravity(false)
			return left
		}
		if rl.CheckCollisionRecs(p.cRects[right].Rect(), s.HitBox().Rect()) &&
			p.actions[wall] {
			p.SetGravity(false)
			return right
		}
	}
	if p.platform != nil {
		if rl.CheckCollisionRecs(p.cRects[floor].Rect(), p.platform.HitBox().Rect()) {
			return floor
		} else {
			p.platform = nil
		}
	}
	p.SetGravity(true)
	return air
}
