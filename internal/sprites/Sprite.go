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
	s := ObjectSprite{
		Pos: newPos(obj, a),
	}
	var err error
	s.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *ObjectSprite) Draw(src rl.Texture2D, pos *Pos) {
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, s.rect.Width*s.FlipH, s.rect.Height*s.FlipV),
		rl.Rectangle(*s.Rect()),
		rl.Vector2{}, 0, rl.White,
	)
}
