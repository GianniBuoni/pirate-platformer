package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	speed     float32
	direction rl.Vector2
	image     string
	rect      rl.Rectangle
}

func NewSprite(image string, pos rl.Vector2) (Sprite) {
  return &BasicSprite{
    image: image,
    rect: rl.Rectangle{
      X: pos.X, Y: pos.Y,
    },
  }
}
