package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	image     string
	rect      SpriteRect
	oldRect   SpriteRect
	direction rl.Vector2
	speed     float32
}

func NewSprite(image string, pos rl.Vector2, a Assets) (Sprite, error) {
	src, err := a.GetImage(ImageLib, image)
	if err != nil {
		return nil, fmt.Errorf(
			"New sprite %s, could not be created. %w", image, err,
		)
	}

	sprite := &BasicSprite{
		image: image,
	}

	sprite.rect = rects.NewRectangle(
		pos.X, pos.Y,
		float32(src.Width), float32(src.Height),
	)

	return sprite, nil
}
