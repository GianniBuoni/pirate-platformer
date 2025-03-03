package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RadialSprite struct {
	ID
	Pos
	RadialMove
}

func NewRadialSprite(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	rs := RadialSprite{
		Pos:        newPos(obj, a),
		RadialMove: newRadial(obj),
	}
	var err error
	rs.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

func (rs *RadialSprite) Update() error {
	rs.moveR(rs.oldRect, rs.hitbox)
	rs.Pos.Update()
	return nil
}

func (rs *RadialSprite) Draw(src rl.Texture2D, pos *Pos) {
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, rs.rect.Width, rs.rect.Height),
		rl.Rectangle(*rs.rect),
		rl.Vector2{}, 0, rl.White,
	)
}
