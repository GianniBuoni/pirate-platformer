package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
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

func NewText(o Object) Text {
	text := Text{
		string: o.Image,
		pos:    rl.NewVector2(o.X, o.Y),
		color:  rl.White,
	}
	switch o.Properties.Loader {
	case "bodyCenter":
		text.alignment = center
	case "displayCenter":
		text.alignment = center
		text.face = display
	}
	return text
}

func (t Text) Draw(string string, font rl.Font) {
	var fontSize float32
	switch t.face {
	case display:
		fontSize = DisplayTextSize
	default:
		fontSize = BodyTextSize
	}

	measure := rl.MeasureTextEx(font, string, fontSize, TextSpacing)

	pos := t.pos
	switch t.alignment {
	case center:
		pos.X -= measure.X / 2
	default:
		pos.Y -= measure.Y / 2
	}

	rl.DrawTextEx(font, string, pos, fontSize, TextSpacing, t.color)
}
