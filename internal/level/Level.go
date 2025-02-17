package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type LevelData struct {
	player      Sprite
	levelAssets Assets
	mapData     *tiled.Map
	groups      map[string][]Sprite
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
		groups:      map[string][]Sprite{},
	}, nil
}

func (l *LevelData) Update() {
	//l.player.Update()
}

func (l *LevelData) Draw() error {
	allSprites, ok := l.groups["all"]
	if ok {
		for _, sprite := range allSprites {
			err := sprite.Draw(l.levelAssets)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *LevelData) PlayerPos() rl.Vector2 {
	return l.player.HitBox().Center()
}
