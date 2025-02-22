package lib

type Sprite interface {
	Update()
	Draw() error
	Rect() *Rect
	HitBox() *Rect
	OldRect() *Rect
}
