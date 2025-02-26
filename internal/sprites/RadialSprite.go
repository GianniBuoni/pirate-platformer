package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RadialSprite struct {
	Pos
	ID
	RadialMove
}

func NewRadialSprite(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	rs := RadialSprite{
		Pos:        newPos(obj, a),
		ID:         id,
		RadialMove: newRadial(obj),
	}
	return &rs, nil
}

func (rs *RadialSprite) Update() {
	rs.moveR(rs.oldRect, rs.hitbox)
	rs.Pos.Update()
}

func (rs *RadialSprite) Draw() error {
	src, err := rs.assets.GetImage(rs.assetLib, rs.image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, rs.rect.Width, rs.rect.Height),
		rl.Rectangle(*rs.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}
