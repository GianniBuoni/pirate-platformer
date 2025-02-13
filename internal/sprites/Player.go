package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerData struct {
	BasicSprite
	hitbox           SpriteRect
	collisionSprites *[]Sprite
	frameCount       int
	frameIndex       float32
	frameSize        float32
	frameSpeed       float32
	gravity          float32
}

func NewPlayer(pos rl.Vector2, a Assets, s *[]Sprite) (Sprite, error) {
	state := "idle"
	src, err := a.GetImage(PlayerLib, state)
	if err != nil {
		return nil, fmt.Errorf(
			"New player with state: %s, could not be created. %w",
			state, err,
		)
	}

	sprite := &PlayerData{
		frameSize:        96,
		frameSpeed:       FrameSpeed,
		gravity:          Gravity,
		collisionSprites: s,
	}
	sprite.image = state
	sprite.frameCount = int(float32(src.Width) / sprite.frameSize)

	sprite.rect = rects.NewRectangle(
		pos.X, pos.Y-sprite.frameSize*2,
		sprite.frameSize*2, sprite.frameSize*2,
	)
	sprite.hitbox = rects.NewRectangle(
		sprite.rect.Left()+TileSize,
		sprite.rect.Top()+TileSize,
		40, TileSize,
	)
	sprite.oldRect = rects.NewRectangle(
		sprite.hitbox.Left(), sprite.hitbox.Top(),
		sprite.hitbox.Rect().Width, sprite.hitbox.Rect().Height,
	)

	return sprite, nil
}
