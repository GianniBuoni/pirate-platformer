package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	groups  map[string][]int
	spirtes map[int]Sprite
	assets  *Assets
	camera  *CameraRect
	player  *Player
	stats   *Stats
	nextID  int
}

func NewLevel(stats *Stats, assets *Assets) *Level {
	l := &Level{
		groups:  map[string][]int{},
		spirtes: map[int]Sprite{},
		assets:  assets,
		stats:   stats,
	}
	return l
}

func (l *Level) Update() error {
	return nil
}

func (l *Level) Assets() *Assets {
	return l.assets
}

func (l *Level) NextId() int {
	id := l.nextID
	l.nextID++
	return id
}

func (l *Level) Sprites() map[int]Sprite {
	return l.spirtes
}

func (l *Level) AddSpriteGroup(
	s Sprite, spriteMap map[int]Sprite, groups ...string,
) {
	spriteMap[s.GetID().GID] = s
	for _, key := range groups {
		l.groups[key] = append(l.groups[key], s.GetID().GID)
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
