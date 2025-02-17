package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *BasicSprite) Update() {}

func (s *BasicSprite) Draw(a Assets) error {
	src, err := a.GetImage(s.assetLib, s.image)
	if err != nil {
		return err
	}
	srcRect := rl.NewRectangle(
		s.imgRect.X, s.imgRect.Y,
		s.imgRect.Width*s.flipH,
		s.imgRect.Height*s.flipV,
	)
	rl.DrawTexturePro(
		src, srcRect, s.rect.Rect(), rl.Vector2{}, 0, rl.White,
	)
	return nil
}
