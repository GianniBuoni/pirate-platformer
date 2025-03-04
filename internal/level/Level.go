package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelData struct {
	MapProps  MapProps `json:"properties"`
	assets    *Assets
	camera    *CameraRect
	player    *Player
	stats     *Stats
	sprites   map[int]Sprite
	groups    map[string][]int
	tileRefs  map[GIDRange]string
	pathRects map[int]*Rect
	nextId    int
	Width     int `json:"width"`
	Height    int `json:"height"`
}

func NewLevel(stats *Stats, assets *Assets) (*LevelData, error) {
	l := LevelData{
		assets:    assets,
		stats:     stats,
		tileRefs:  map[GIDRange]string{},
		sprites:   map[int]Sprite{},
		groups:    map[string][]int{},
		pathRects: map[int]*Rect{},
	}

	return &l, nil
}

func (l *LevelData) Draw() {
	allSprites, ok := l.groups["all"]
	if ok {
		for _, id := range allSprites {
			s := l.sprites[id]
			s.Draw(s.GetID().Src, s.GetPos())
		}
	}
}

func (l *LevelData) Assets() *Assets {
	return l.assets
}

func (l *LevelData) NextId() int {
	id := l.nextId
	l.nextId++
	return id
}

func (l *LevelData) Sprites() map[int]Sprite {
	return l.sprites
}

func (l *LevelData) Texts() map[string]Text {
	return map[string]Text{}
}

func (l *LevelData) CameraPos() rl.Vector2 {
	if l.player == nil {
		return rl.NewVector2(WindowW/2, WindowH/2)
	}
	return l.camera.CamTarget
}
