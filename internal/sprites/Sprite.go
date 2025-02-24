package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectSprite struct {
	ID
	Pos
}

func NewSprite(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	s := ObjectSprite{
		ID:  id,
		Pos: newPos(obj, a),
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
		rl.NewRectangle(0, 0, s.rect.Width*s.flipH, s.rect.Height*s.flipV),
		rl.Rectangle(*s.Rect()),
		rl.Vector2{}, 0, rl.White,
	)
	//drawRect(s.hitbox, rl.Blue)
	return nil
}
