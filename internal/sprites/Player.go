package sprites

import (
	"fmt"
	"sync"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
)

type PlayerData struct {
	BasicSprite
	mu               sync.RWMutex
	hitbox           SpriteRect
	collisionSprites *[]Sprite
	platformSprites  *[]Sprite
	actions          map[PlayerState]bool
	frameCount       int
	frameIndex       float32
	frameSize        float32
	frameSpeed       float32
	gravity          float32
	hitboxOffset     float32
}

func NewPlayer(args NewPlayerParams) (Sprite, error) {
	state := "idle"
	src, err := args.Assets.GetImage(PlayerLib, state)
	if err != nil {
		return nil, fmt.Errorf(
			"New player with state: %s, could not be created. %w",
			state, err,
		)
	}
	// INIT CONSTANTS AND ARGUMENTS
	p := &PlayerData{
		collisionSprites: args.CSprites,
		platformSprites:  args.PSprites,
		frameSize:        96,
		frameSpeed:       FrameSpeed,
		gravity:          Gravity,
	}

	// INIT ANIMATION DATA
	p.frameCount = int(float32(src.Width) / p.frameSize)
	p.speed = PlayerSpeed
	p.flip = 1

	// INIT ACTION MAP
	p.actions = map[PlayerState]bool{
		attack:      false,
		canAttack:   true,
		canPlatform: true,
		fall:        false,
		run:         true,
		wall:        false,
	}

	// INIT RECT DATA
	p.rect = rects.NewRectangle(
		args.Pos.X, args.Pos.Y,
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
