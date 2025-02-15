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

func (s *BasicSprite) OldRect() SpriteRect {
	return s.oldRect
}

func (s *PlayerData) HitBox() SpriteRect {
	return s.hitbox
}

func (s *PlayerData) IsAttacking() bool {
	return s.actions[attack]
}
