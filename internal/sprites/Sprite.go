package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectSprite struct {
	ID
	Pos
}

func NewSprite(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	id, err := newId(obj, aLib, a)
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

func (s *ObjectSprite) Draw(id *ID, pos *Pos) error {
	src, err := s.assets.GetImage(s.assetLib, s.Image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, s.rect.Width*s.FlipH, s.rect.Height*s.FlipV),
		rl.Rectangle(*s.Rect()),
		rl.Vector2{}, 0, rl.White,
	)
	//drawRect(s.hitbox, rl.Blue)
	return nil
}
