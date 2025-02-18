package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
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

func (s *BasicSprite) FlipV() {
	s.flipV *= -1
}

func (s *BasicSprite) FlipH() {
	s.flipH *= -1
}

func (s *PlayerData) IsAttacking() bool {
	return s.actions[attack]
}
