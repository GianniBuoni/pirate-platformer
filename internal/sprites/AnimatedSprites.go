package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	BasicSprite
	frameIndex float32
	frameSpeed float32
	frameSize  float32
	frameCount int
}

func NewAnimatedSprite(
	image string, pos rl.Vector2, a Assets, opts ...func(*BasicSprite),
) (*AnimatedSprite, error) {
	as := &AnimatedSprite{}
	return as, nil
}
