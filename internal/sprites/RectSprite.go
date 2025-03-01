package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectSprite struct {
	Pos
	color rl.Color
}

func NewRectSrite(
	obj Object, color rl.Color, a *Assets,
) (Sprite, error) {
	rs := &RectSprite{
		Pos:   newPos(obj, a),
		color: color,
	}
	return rs, nil
}

func (rs *RectSprite) Draw(*ID, *Pos) error {
	rl.DrawRectangleRec(
		rl.Rectangle(*rs.rect), rs.color,
	)
	return nil
}

func (rs *RectSprite) Update() {}
func (rs *RectSprite) GetID() *ID {
	return &ID{}
}
