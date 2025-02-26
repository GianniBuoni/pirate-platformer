package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Sprite interface {
	Update()
	Draw() error
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
}
