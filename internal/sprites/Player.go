package sprites

import (
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
	mu       *sync.RWMutex
	platform Sprite
	cRects   map[CollisionSide]*Rect
	groups   map[string][]Sprite
	actions  map[PlayerState]bool
	gravity  float32
}

func NewPlayer(args NewPlayerParams) (*PlayerData, error) {
	requiredGroups := []string{"collision", "platform", "damage"}
	for _, group := range requiredGroups {
		if _, ok := args.Groups[group]; !ok {
			return nil,
				fmt.Errorf("Error with player init: %s sprite group is nil", group)
		}
	}
	p := initalPlayer
	p.groups = args.Groups

	// CALC INITAL RECT DATA
	p.rect = NewRectangle(
		args.Pos.X, args.Pos.Y,
		p.imgRect.Width*2, p.imgRect.Width*2,
	)
	p.hitbox = NewRectangle(
		p.rect.X+p.hitboxOffset.X,
		p.rect.Y+p.hitboxOffset.Y,
		hitboxW, TileSize,
	)
	p.oldRect.Copy(p.hitbox)

	p.cRects[floor] = NewRectangle(
		p.hitbox.Left(), p.hitbox.Bottom(),
		p.hitbox.Rect().Width, 10,
	)
	p.cRects[left] = NewRectangle(
		p.hitbox.Left()-2, p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	p.cRects[right] = NewRectangle(
		p.hitbox.Right(), p.hitbox.Top()+2, 2, p.hitbox.Rect().Height/2,
	)
	return &p, nil
}
