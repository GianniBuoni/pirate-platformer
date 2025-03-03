package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	font      rl.Font
	string    string
	pos       rl.Vector2
	alignment textAlignment
	face      textFace
	color     rl.Color
}

type textAlignment uint
type textFace uint

const (
	left textAlignment = iota
	right
	center
)

const (
	body textFace = iota
	display
)

func (t Text) Draw(string string) {
	var fontSize float32
	switch t.face {
	case display:
		fontSize = DispalyTextSize
	default:
		fontSize = BodyTextSize
	}

	measure := rl.MeasureTextEx(t.font, string, fontSize, TextSpacing)

	var pos rl.Vector2
	switch t.alignment {
	default:
		pos = rl.NewVector2(t.pos.X, t.pos.Y-(measure.Y/2))
	}

	rl.DrawTextEx(t.font, string, pos, fontSize, TextSpacing, t.color)
}
