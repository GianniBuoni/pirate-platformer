package sprites

import (
	"errors"
	"fmt"
	"sync"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NewPlayerParams struct {
	Assets Assets
	Pos    rl.Vector2
	Groups map[string][]Sprite
}

type PlayerData struct {
	AnimatedSprite
	mu      sync.RWMutex
	groups  map[string][]Sprite
	actions map[PlayerState]bool
	gravity float32
}

func NewPlayer(args NewPlayerParams) (*PlayerData, error) {
	if _, ok := args.Groups["collision"]; !ok {
		return nil, errors.New(
			"error making new player: collision sprite group is nil",
		)
	}
	if _, ok := args.Groups["platform"]; !ok {
		return nil, errors.New(
			"error making new player: platform sprite group is nil",
		)
	}
	if _, ok := args.Groups["damage"]; !ok {
		return nil, errors.New(
			"error making new player: damage sprite group is nil",
		)
	}
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
		gravity: Gravity,
		groups:  args.Groups,
	}

	// INIT ANIMATION DATA
	p.imgRect = rl.NewRectangle(0, 0, 96, 96)
	p.frameSpeed = FrameSpeed
	p.frameCount = int(float32(src.Width) / p.imgRect.Width)
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
		p.imgRect.Width*2, p.imgRect.Width*2,
	)
	var hitboxW float32
	hitboxW = 48
	p.hitboxOffset = rl.NewVector2(
		TileSize+(TileSize-hitboxW)/2,
		TileSize,
	)
	p.hitbox = NewRectangle(
		p.rect.X+p.hitboxOffset.X,
		p.rect.Y+p.hitboxOffset.Y,
		hitboxW, TileSize,
	)
	p.oldRect = &Rect{}
	p.oldRect.Copy(p.hitbox)
	return p, nil
}
