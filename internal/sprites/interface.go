package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Sprite interface {
	Update()
	Draw(*ID, *Pos) error
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
	GetID() *ID
	GetPos() *Pos
}
