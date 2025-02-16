package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	BasicSprite
	hitbox       SpriteRect
	frameIndex   float32
	frameSpeed   float32
	frameSize    float32
	frameCount   int
	hitboxOffset float32
}

func NewAnimatedSprite(
	image string, pos rl.Vector2, a Assets, flipH, flipV bool,
) (*AnimatedSprite, error) {
	src, err := a.GetImage(ImageLib, image)
	if err != nil {
		return nil, fmt.Errorf(
			"New sprite with image: %s, could not be created. %w",
			image, err,
		)
	}
	as := &AnimatedSprite{
		frameSpeed: FrameSpeed,
		frameSize:  TileSize, // TODO find a dynamic way to set this field
		frameCount: int(float32(src.Width) / TileSize),
	}
	as.image = image
	as.rect = rects.NewRectangle(
		pos.X, pos.Y, as.frameSize, float32(src.Height),
	)
	as.oldRect = rects.NewRectangle(
		pos.X, pos.Y, as.frameSize, float32(src.Height),
	)
	if flipH {
		as.flipH = -1
	} else {
		as.flipH = 1
	}
	if flipV {
		as.flipH = -1
	} else {
		as.flipV = 1
	}
	return as, nil
}
