package level

import (
	"encoding/json"
	"os"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type LevelData struct {
	Layers      []Layer `json:"layers"`
	tiles       []*tiled.Layer
	levelAssets *assets.Assets
	player      Sprite
	groups      map[string][]Sprite
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
}

func NewLevel(assets *assets.Assets, mapPath string) (*LevelData, error) {
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}
	tiledData, err := tiled.LoadFile(mapPath)
	if err != nil {
		return nil, err
	}
	l := LevelData{levelAssets: assets}

	json.Unmarshal(data, &l)
	l.tiles = tiledData.Layers
	return &l, nil
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
