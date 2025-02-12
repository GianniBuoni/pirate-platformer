package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	BasicSprite
	hitbox     rl.Rectangle
	frameCount int
	frameIndex float32
	frameSize  float32
	frameSpeed float32
	gravity    float32
}

func NewPlayer(pos rl.Vector2, a Assets) (Sprite, error) {
	state := "run"
	src, err := a.GetImage(PlayerLib, state)
	if err != nil {
		return nil, fmt.Errorf(
			"New player with state: %s, could not be created. %w",
			state, err,
		)
	}

	sprite := &Player{
		frameSize:  96,
		frameSpeed: lib.FrameSpeed,
	}
	sprite.image = state
	sprite.frameCount = int(float32(src.Width) / sprite.frameSize)

	rect := rl.NewRectangle(
		pos.X, pos.Y-sprite.frameSize*2,
		sprite.frameSize*2, sprite.frameSize*2,
	)
	sprite.rect = rect
	return sprite, nil
}
