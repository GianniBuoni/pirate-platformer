package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/lafriks/go-tiled"
)

type LevelData struct {
	allSprites  []Sprite
	player      Sprite
	levelAssets Assets
	mapData     *tiled.Map
}

func NewLevel(assets Assets, mapPath string) (Level, error) {
	mapData, err := tiled.LoadFile(mapPath)
	if err != nil {
		return nil,
			fmt.Errorf("error loading map: %s, %w", mapPath, err)
	}
	return &LevelData{
		levelAssets: assets,
		mapData:     mapData,
	}, nil
}

func (l *LevelData) Update() {}

func (l *LevelData) Draw() error {
	if len(l.allSprites) != 0 {
		for _, sprite := range l.allSprites {
			err := sprite.Draw(l.levelAssets)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
