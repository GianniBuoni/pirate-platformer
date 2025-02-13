package interfaces

import rl "github.com/gen2brain/raylib-go/raylib"

type Sprite interface {
	Update()
	Draw(Assets) error

	Rect() rl.Rectangle
	Top() float32
	Bottom() float32
	Left() float32
	Right() float32

	OffsetCentre()
}
