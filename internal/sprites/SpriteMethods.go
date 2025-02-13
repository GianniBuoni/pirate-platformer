package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *BasicSprite) Update() {}

func (s *BasicSprite) Draw(a Assets) error {
	src, err := a.GetImage(ImageLib, s.image)
	if err != nil {
		return err
	}

	rl.DrawTextureV(
		src,
		rl.Vector2{X: s.rect.Left(), Y: s.rect.Right()},
		rl.White,
	)
	return nil
}
