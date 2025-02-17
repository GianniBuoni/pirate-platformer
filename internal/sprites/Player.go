package sprites

import (
	"fmt"
	"sync"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NewPlayerParams struct {
	Pos      rl.Vector2
	Assets   Assets
	CSprites *[]Sprite
	PSprites *[]Sprite
}

type PlayerData struct {
	AnimatedSprite
	mu               sync.RWMutex
	collisionSprites *[]Sprite
	platformSprites  *[]Sprite
	actions          map[PlayerState]bool
	gravity          float32
}

func NewPlayer(args NewPlayerParams) (*PlayerData, error) {
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
		gravity:          Gravity,
	}

	// INIT ANIMATION DATA
	p.frameSize = 96
	p.frameSpeed = FrameSpeed
	p.frameCount = int(float32(src.Width) / p.frameSize)
	p.speed = PlayerSpeed
	p.flipH = 1

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
	p.rect = NewRectangle(
		args.Pos.X, args.Pos.Y,
		p.frameSize*2, p.frameSize*2,
	)
	var hitboxW float32
	hitboxW = 48
	p.hitboxOffset = TileSize + (TileSize-hitboxW)/2
	p.hitbox = NewRectangle(
		p.rect.Left()+p.hitboxOffset,
		p.rect.Top()+TileSize,
		hitboxW, TileSize,
	)
	p.oldRect = NewRectangle(
		p.hitbox.Left(), p.hitbox.Top(),
		p.hitbox.Rect().Width, p.hitbox.Rect().Height,
	)
	return p, nil
}
