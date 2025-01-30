package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	frameIndex  float32
	frameOffset float32
	frameSpeed  float32
	frameCount  int
	posOffset   rl.Vector2
	hitbox      rl.Rectangle
	BasicSprite
}

func NewPlayer(pos rl.Vector2, a Assets) (Sprite, error) {
	state := "run"
	src, err := a.GetPlayer(state)
	if err != nil {
		return nil, fmt.Errorf(
			"New player with state: %s, could not be created. %w",
			state, err,
		)
	}

	sprite := &Player{
		frameOffset: 96,
		frameSpeed:  lib.FrameSpeed,
		posOffset:   rl.Vector2{X: 32, Y: 32},
	}
	sprite.image = state
	sprite.frameCount = int(float32(src.Width) / sprite.frameOffset)

	rect := rl.NewRectangle(
		pos.X, pos.Y,
		sprite.frameOffset*2, sprite.frameOffset*2,
	)
	sprite.rect = rect
	return sprite, nil
}
