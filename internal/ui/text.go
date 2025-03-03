package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	string    string
	pos       rl.Vector2
	alignment textAlignment
	face      textFace
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
