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

func (p *Player) checkCollisionSide() CollisionSide {
	for _, s := range p.Groups["collision"] {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.cRects[floor]), rl.Rectangle(*s.HitBox()),
		) {
			return floor
		}
	}
	for _, s := range p.Groups["wall"] {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.cRects[left]), rl.Rectangle(*s.HitBox()),
		) {
			return left
		}
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.cRects[right]), rl.Rectangle(*s.HitBox()),
		) {
			return right
		}
	}
	for _, s := range p.Groups["platform"] {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.cRects[floor]), rl.Rectangle(*s.HitBox()),
		) {
			return floor
		}
	}
	return air
}

func (p *Player) getCRects() {
	p.cRects[floor] = NewRectangle(0, 0, p.hitbox.Width, 2)
	p.cRects[left] = NewRectangle(0, 0, 2, p.hitbox.Height-4)
	p.cRects[right] = NewRectangle(0, 0, 2, p.hitbox.Height-4)
}

func (p *Player) updateCRects() {
	p.cRects[floor].Set(
		Center(p.hitbox.Center().X, p.hitbox.Center().Y), Top(p.hitbox.Bottom()),
	)
	p.cRects[left].Set(
		Bottom(p.hitbox.Bottom()), Right(p.hitbox.Left()),
	)
	p.cRects[right].Set(
		Bottom(p.hitbox.Bottom()), Left(p.hitbox.Right()),
	)
}
