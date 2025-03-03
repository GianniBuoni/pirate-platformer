package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite interface {
	Update() error
	Draw(rl.Texture2D, *Pos)

	// Pos
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
	GetPos() *Pos

	// ID
	GetID() *ID
}
