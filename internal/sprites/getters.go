package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
)

func (s *BasicSprite) Rect() SpriteRect {
	return s.rect
}

func (s *BasicSprite) HitBox() SpriteRect {
	return s.rect
}

func (s *PlayerData) HitBox() SpriteRect {
	return s.hitbox
}
