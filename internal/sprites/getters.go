package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

func (s *BasicSprite) Pos() rl.Vector2 {
	return rl.NewVector2(s.rect.X, s.rect.Y)
}
