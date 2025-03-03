package level

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelData struct {
	MapProps    MapProps `json:"properties"`
	Layers      []Layer  `json:"layers"`
	levelAssets *Assets
	camera      *CameraRect
	player      *Player
	stats       *Stats
	groups      map[string][]Sprite
	tileRefs    map[GIDRange]string
	pathRects   map[int]*Rect
	Width       int `json:"width"`
	Height      int `json:"height"`
}

func NewLevel(
	stats *Stats, assets *Assets, mapPath string,
) (*LevelData, error) {
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	l := LevelData{
		levelAssets: assets,
		stats:       stats,
		tileRefs:    map[GIDRange]string{},
		groups:      map[string][]Sprite{},
		pathRects:   map[int]*Rect{},
	}
	t := TileRefs{}

	json.Unmarshal(data, &l)
	json.Unmarshal(data, &t)

	for _, ref := range t.TileRef {
		name := GetAssetKey(ref.Source)
		tileset, ok := l.levelAssets.TilesetData[name]
		if !ok {
			return nil,
				fmt.Errorf("key: %s\n not found in levelAssets tileset data.", name)
		}
		key := GIDRange{
			FirstGID: ref.FirstGID,
			LastGID:  ref.FirstGID + tileset.Count - 1,
		}
		l.tileRefs[key] = name
	}
	return &l, nil
}

func (l *LevelData) Draw() {
	allSprites, ok := l.groups["all"]
	if ok {
		for _, s := range allSprites {
			s.Draw(s.GetID().Src, s.GetPos())
		}
	}
}

func (l *LevelData) CameraPos() rl.Vector2 {
	if l.player == nil {
		return rl.NewVector2(WindowW/2, WindowH/2)
	}
	return l.camera.CamTarget
}
