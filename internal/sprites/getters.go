package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *BasicSprite) Rect() *Rect {
	return s.rect
}

func (s *BasicSprite) HitBox() *Rect {
	return s.hitbox
}

func (s *BasicSprite) OldRect() *Rect {
	return s.oldRect
}

func (s *BasicSprite) Movement(r *Rect) bool {
	dt := rl.GetFrameTime()
	r.X += s.direction.X * s.speed * dt
	r.Y += s.direction.Y * s.speed * dt
	if s.direction.Y != 0 {
		return true
	}
	return false
}

func (s *BasicSprite) FlipV() {
	s.flipV *= -1
}

func (s *BasicSprite) FlipH() {
	s.flipH *= -1
}

func (s *BasicSprite) SetHitbox(
	offset rl.Vector2, width, height float32,
) {
	s.hitbox = NewRectangle(
		s.rect.X+offset.X, s.rect.Y+offset.Y,
		width, height,
	)
	s.oldRect = &Rect{}
	s.oldRect.Copy(s.hitbox)
}

func (s *PlayerData) IsAttacking() bool {
	return s.actions[attack]
}
