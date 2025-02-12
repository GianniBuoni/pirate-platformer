package level

import (
	"strings"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *LevelData) Load() error {
	err := l.loadTiles()
	if err != nil {
		return err
	}
	err = l.loadPlayer()
	if err != nil {
		return err
	}
	return nil
}

func (l *LevelData) loadPlayer() error {
	for _, objGroup := range l.mapData.ObjectGroups {
		if objGroup.Name == "Objects" {
			for _, obj := range objGroup.Objects {
				if obj.Name == "player" {
					sprite, err := sprites.NewPlayer(
						rl.Vector2{X: float32(obj.X), Y: float32(obj.Y)},
						l.levelAssets,
					)
					if err != nil {
						return err
					}
					l.AddPlayer(sprite)
				}
			}
		}
	}
	return nil
}

func (l *LevelData) loadTiles() error {
	for _, layer := range l.mapData.Layers {
		for i, tile := range layer.Tiles {
			if !tile.IsNil() {
				tilePos := rl.Vector2{
					X: float32(i%l.mapData.Width) * TileSize,
					Y: float32(i/l.mapData.Width) * TileSize,
				}
				srcRect := tile.GetTileRect()
				srcPos := rl.Vector2{
					X: float32(srcRect.Min.X),
					Y: float32(srcRect.Min.Y),
				}
				srcKey := GetAssetKey(tile.Tileset.Image.Source)

				sprite, err := sprites.NewTileSprite(
					srcKey, tilePos, srcPos, l.levelAssets,
				)
				if err != nil {
					return err
				}

				if strings.Contains(layer.Name, "BG") ||
					strings.Contains(layer.Name, "FG") {
					l.AddSpriteGroup(sprite, "all")
				} else {
					l.AddSpriteGroup(sprite, "all", "collision")
				}
			}
		}
	}
	return nil
}
