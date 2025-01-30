package sprites

import (
	"fmt"
	"log"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	frameIndex  float32
	frameOffset float32
	frameSpeed  float32
	posOffset   rl.Vector2
	hitbox      rl.Rectangle
	BasicSprite
}

func NewPlayer(pos rl.Vector2, a Assets) (Sprite, error) {
	state := "idle"
	_, err := a.GetPlayer(state)
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

	rect := rl.NewRectangle(
		pos.X, pos.Y,
		sprite.frameOffset*2, sprite.frameOffset*2,
	)
	sprite.rect = rect
	return sprite, nil
}

func (p *Player) Draw(a Assets) {
	src, err := a.GetPlayer(p.image)
	if err != nil {
		log.Fatal(err)
	}

	srcRect := rl.NewRectangle(
		p.frameOffset*p.frameIndex,
		p.frameOffset*p.frameIndex,
		p.frameOffset,
		p.frameOffset,
	)

	rl.DrawTexturePro(
		src,
		srcRect,
		p.rect,
		rl.Vector2Scale(p.posOffset, 2),
		0,
		rl.White,
	)
}
