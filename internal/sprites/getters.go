package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

func (s *BasicSprite) Pos() rl.Vector2 {
	return rl.NewVector2(s.rect.X, s.rect.Y)
}

func (s *BasicSprite) Rect() rl.Rectangle {
	return s.rect
}

func (s *BasicSprite) Top() float32 {
	return s.rect.Y
}

func (s *BasicSprite) Bottom() float32 {
	return s.rect.Y + s.rect.Height
}

func (s *BasicSprite) Left() float32 {
	return s.rect.X + s.rect.Width
}

func (s *BasicSprite) Right() float32 {
	return s.rect.X + s.rect.Width
}

func (s *PlayerData) Top() float32 {
	return s.hitbox.Y
}

func (s *PlayerData) Bottom() float32 {
	return s.hitbox.Y + s.hitbox.Height
}

func (s *PlayerData) Left() float32 {
	return s.hitbox.X
}

func (s *PlayerData) Right() float32 {
	return s.hitbox.X + s.hitbox.Width
}
