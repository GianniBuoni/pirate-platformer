package interfaces

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Sprite interface {
	Update()
	Draw(Assets) error

	// returns the SpriteRect interface
	Rect() *Rect
	OldRect() *Rect
	Movement(*Rect) bool
	// returns an an underlying raylib Rectangle
	HitBox() *Rect
}
