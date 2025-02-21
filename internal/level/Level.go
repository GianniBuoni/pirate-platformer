package level

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelData struct {
	Layers      []Layer `json:"layers"`
	levelAssets *assets.Assets
	player      Sprite
	groups      map[string][]Sprite
	tileRefs    map[GIDRange]string
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
}

func NewLevel(assets *assets.Assets, mapPath string) (*LevelData, error) {
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	l := LevelData{
		levelAssets: assets,
		tileRefs:    map[GIDRange]string{},
		groups:      map[string][]Sprite{}}
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
	fmt.Printf("l: %v\n", l)
	return &l, nil
}

func (l *LevelData) GetImageRect(gid int) (x, y float32, err error) {
	var (
		name     string
		firstGID int
	)

	for k, v := range l.tileRefs {
		if k.FirstGID <= gid && k.LastGID >= gid {
			name = v
			firstGID = k.FirstGID
		}
	}
	if name == "" {
		return 0, 0, fmt.Errorf("gid: %d\n not found in tileset refs.", gid)
	}
	tileset := l.levelAssets.TilesetData[name]

	idx := firstGID - gid
	if idx > tileset.Count {
		fmt.Printf("index: %d out of range of tileset '%s'.\n", idx, name)
		return 0, 0,
			errors.New("Check if there are rotation flags on tile gid")
	}
	x = float32(idx%tileset.Columns) * TileSize
	y = float32(idx/tileset.Columns) * TileSize
	return x, y, nil
}

func (l *LevelData) Update() {
	/*
		for _, mSprite := range l.groups["moving"] {
			mSprite.Update()
		}
		l.player.Update()
	*/
}

func (l *LevelData) Draw() error {
	/*
		allSprites, ok := l.groups["all"]
		if ok {
			for _, sprite := range allSprites {
				err := sprite.Draw(l.levelAssets)
				if err != nil {
					return err
				}
			}
		}
	*/
	return nil
}

func (l *LevelData) PlayerPos() rl.Vector2 {
	return l.player.HitBox().Center()
}
