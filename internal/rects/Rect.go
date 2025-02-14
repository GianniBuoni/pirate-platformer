package rects

import rl "github.com/gen2brain/raylib-go/raylib"

type Rect rl.Rectangle

func NewRectangle(x, y, width, height float32) *Rect {
	return &Rect{
		X: x, Y: y, Width: width, Height: height,
	}
}

func (r *Rect) Rect() rl.Rectangle {
	return rl.NewRectangle(
		r.X, r.Y, r.Width, r.Height,
	)
}

func (r *Rect) Top() float32 {
	return r.Y
}

func (r *Rect) Bottom() float32 {
	return r.Y + r.Height
}

func (r *Rect) Left() float32 {
	return r.X
}

func (r *Rect) Right() float32 {
	return r.X + r.Width
}

func (r *Rect) Center() rl.Vector2 {
	return rl.NewVector2(
		r.X+(r.Width/2),
		r.Y+(r.Height/2),
	)
}

func (r *Rect) Set(opts ...func(*Rect)) {
	new := Rect{
		X: r.X, Y: r.Y, Width: r.Width, Height: r.Height,
	}
	for _, opt := range opts {
		opt(&new)
	}
	if r.X != new.X {
		r.X = new.X
	}
	if r.Y != new.Y {
		r.Y = new.Y
	}
	if r.Width != new.Width {
		r.Width = new.Width
	}
	if r.Height != new.Height {
		r.Height = new.Height
	}
}

func Top(y float32) func(*Rect) {
	return func(r *Rect) {
		r.Y = y
	}
}

func Bottom(y float32) func(*Rect) {
	return func(r *Rect) {
		r.Y = y - r.Height
	}
}

func Left(x float32) func(*Rect) {
	return func(r *Rect) {
		r.X = x
	}
}

func Right(x float32) func(*Rect) {
	return func(r *Rect) {
		r.X = x - r.Width
	}
}
