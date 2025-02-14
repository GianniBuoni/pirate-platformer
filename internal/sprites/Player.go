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
	actions          map[string]bool
	frameCount       int
	frameIndex       float32
	frameSize        float32
	frameSpeed       float32
	gravity          float32
	hitboxOffset     float32
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
	// INIT CONSTANTS AND ARGUMENTS
	p := &PlayerData{
		collisionSprites: s,
		frameSize:        96,
		frameSpeed:       FrameSpeed,
		gravity:          Gravity,
	}

	// INIT ANIMATION DATA
	p.frameCount = int(float32(src.Width) / p.frameSize)
	p.speed = PlayerSpeed
	p.flip = 1

	// INIT ACTION MAP
	p.actions = map[string]bool{
		"wall":   false,
		"attack": false,
	}

	// INIT RECT DATA
	p.rect = rects.NewRectangle(
		pos.X, pos.Y-p.frameSize*2,
		p.frameSize*2, p.frameSize*2,
	)
	var hitboxW float32
	hitboxW = 48
	p.hitboxOffset = TileSize + (TileSize-hitboxW)/2
	p.hitbox = rects.NewRectangle(
		p.rect.Left()+p.hitboxOffset,
		p.rect.Top()+TileSize,
		hitboxW, TileSize,
	)
	p.oldRect = rects.NewRectangle(
		p.hitbox.Left(), p.hitbox.Top(),
		p.hitbox.Rect().Width, p.hitbox.Rect().Height,
	)
	return p, nil
}
