package sprites

import (
	"sync"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This is the basic player data that doesn't need to be calced.
var initalPlayer = PlayerData{
	AnimatedSprite: AnimatedSprite{
		BasicSprite: BasicSprite{
			imgRect:      rl.NewRectangle(0, 0, 96, 96),
			oldRect:      &Rect{},
			speed:        PlayerSpeed,
			flipH:        1,
			hitboxOffset: rl.NewVector2(TileSize+(TileSize-hitboxW)/2, TileSize),
		},
		frameSpeed: FrameSpeed,
	},
	mu:      &sync.RWMutex{},
	cRects:  map[CollisionSide]*Rect{},
	actions: defaultPlayerSates,
	gravity: Gravity,
}

var defaultPlayerSates = map[PlayerState]bool{
	attack:      false,
	canAttack:   true,
	canPlatform: true,
	run:         true,
	wall:        false,
}

var hitboxW float32 = 48
