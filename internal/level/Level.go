package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	groups SpriteGroup
	paths  map[int]*Rect
	assets *Assets
	camera *CameraRect
	player *Player
	stats  *Stats
	nextID int
	Top    float32
	Width  float32
	Height float32
}

func NewLevel(stats *Stats, assets *Assets) *Level {
	l := &Level{
		groups: SpriteGroup{
			IDs:     map[string][]int{},
			Sprites: map[int]Sprite{},
		},
		paths:  map[int]*Rect{},
		assets: assets,
		stats:  stats,
	}
	return l
}

func (l *Level) Assets() *Assets {
	return l.assets
}

func (l *Level) NextId() int {
	id := l.nextID
	l.nextID++
	return id
}

func (l *Level) Sprites() SpriteGroup {
	return l.groups
}

func (l *Level) AddSpriteGroup(s Sprite, groups ...string) {
	// assign id to sprite
	id := l.NextId()
	s.GetID().GID = id
	// add sprites to map id -> Sprite
	l.groups.Sprites[id] = s
	// add id to all passed in group names
	for _, key := range groups {
		l.groups.IDs[key] = append(l.groups.IDs[key], id)
	}
}

func (l *Level) CameraPos() rl.Vector2 {
	if l.player == nil {
		return rl.NewVector2(WindowW/2, WindowH/2)
	}
	return l.camera.CamTarget
}

func (l *Level) Texts() map[string]Text {
	return map[string]Text{}
}
