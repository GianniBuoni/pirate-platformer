package interfaces

import (
	rect "github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite interface {
	Update()
	Draw(Assets) error

	// returns the SpriteRect interface
	Rect() SpriteRect
	// returns an an underlying raylib Rectangle
	HitBox() SpriteRect
}

type SpriteRect interface {
	// returns an underlying raylib Rectangle
	Rect() rl.Rectangle
	Top() float32
	Bottom() float32
	Left() float32
	Right() float32
	Center() rl.Vector2

	// setter to manipulate underlying Raylib rectagle
	Set(...func(*rect.Rect))
}
