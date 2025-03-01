package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectSprite struct {
	Pos
	color       rl.Color
	borderColor rl.Color
}

func NewRectSrite(
	obj Object, color, border rl.Color, a *Assets,
) (Sprite, error) {
	rs := &RectSprite{
		Pos:         newPos(obj, a),
		color:       color,
		borderColor: border,
	}
	return rs, nil
}

func (rs *RectSprite) Draw(*ID, *Pos) error {
	rl.DrawRectangleRec(
		rl.Rectangle(*rs.rect), rs.color,
	)
	if rs.borderColor != (rl.Color{}) {
		rl.DrawLineEx(
			rl.NewVector2(rs.rect.X, rs.rect.Y),
			rl.NewVector2(rs.rect.Right(), rs.rect.Y),
			2, rs.borderColor,
		)
	}
	return nil
}

func (rs *RectSprite) Update() {}
func (rs *RectSprite) GetID() *ID {
	return &ID{}
}
