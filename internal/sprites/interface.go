package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Sprite interface {
	Update()
	Draw(*ID, *Pos) error

	// Pos
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
	GetPos() *Pos

	// ID
	GetID() *ID
}
