package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *BasicSprite) Update() {}

func (s *BasicSprite) Draw(a assets.Assets) error {
	src, err := a.GetImage(s.assetLib, s.image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src, s.rect.Rect(), s.rect.Rect(), rl.Vector2{}, 0, rl.White,
	)
	//s.drawRects(rl.Blue)
	return nil
}

// call this func in the draw method to debug any hitbox and rect issues
func (s *BasicSprite) drawRects(c rl.Color) {
	rl.DrawRectangleRec(s.hitbox.Rect(), rl.ColorAlpha(c, .5))
}
