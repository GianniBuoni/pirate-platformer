package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Sprite interface {
	Update()
	Draw(*ID, *Pos) error

	// Pos
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
	Facing() float32
	GetPos() *Pos

	// ID
	GetID() *ID
	Name() string
	Kill()
	GetKill() bool
}
