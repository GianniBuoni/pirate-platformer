package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectSprite struct {
	ID
	Pos
}

func NewSprite(obj Object, a *assets.Assets) (Sprite, error) {
	id, err := newId(obj.Image, assets.ImageLib, a)
	if err != nil {
		return nil, err
	}
	s := ObjectSprite{
		ID:  id,
		Pos: newPos(obj.X, obj.Y, obj.Width, obj.Height),
	}
	return &s, nil
}

func (s *ObjectSprite) Update() {}

func (s *ObjectSprite) Draw() error {
	src, err := s.assets.GetImage(s.assetLib, s.image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, s.rect.Width, s.rect.Height),
		rl.Rectangle(*s.Rect()),
		rl.Vector2{}, 0, rl.White,
	)
	//s.pos.drawRects(rl.Blue)
	return nil
}
