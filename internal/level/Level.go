package level

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelData struct {
	Layers      []Layer `json:"layers"`
	levelAssets *Assets
	player      Sprite
	groups      map[string][]Sprite
	tileRefs    map[GIDRange]string
	Width       int `json:"width"`
	Height      int `json:"height"`
}

func NewLevel(assets *Assets, mapPath string) (*LevelData, error) {
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
	return &l, nil
}

func (l *LevelData) Update() {
	/*
		for _, mSprite := range l.groups["moving"] {
			mSprite.Update()
		}
	*/
	l.player.Update()
}

func (l *LevelData) Draw() error {
	allSprites, ok := l.groups["all"]
	if ok {
		for _, sprite := range allSprites {
			err := sprite.Draw()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *LevelData) PlayerPos() rl.Vector2 {
	if l.player == nil {
		return rl.NewVector2(WindowW/2, WindowH/2)
	}
	return l.player.HitBox().Center()
}
