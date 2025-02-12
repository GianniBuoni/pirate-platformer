package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	speed     float32
	direction rl.Vector2
	image     string
	rect      rl.Rectangle
}

func NewSprite(image string, pos rl.Vector2, a Assets) (Sprite, error) {
	src, err := a.GetImage(ImageLib, image)
	if err != nil {
		return nil, fmt.Errorf(
			"New sprite %s, could not be created. %w", image, err,
		)
	}

	rect := rl.NewRectangle(
		pos.X, pos.Y,
		float32(src.Width), float32(src.Height),
	)

	sprite := &BasicSprite{
		image: image,
		rect:  rect,
	}

	return sprite, nil
}
