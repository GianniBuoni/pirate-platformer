package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	err := l.drawMap()
	if err != nil {
		return err
	}
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

func (l *LevelData) drawMap() error {
	for _, layer := range l.mapData.Layers {
		for i, tile := range layer.Tiles {
			if !tile.IsNil() {
				srcData := tile.GetTileRect()
				srcRect := rl.Rectangle{
					X:      float32(srcData.Min.X),
					Y:      float32(srcData.Min.Y),
					Width:  TileSize,
					Height: TileSize}
				destRect := rl.Rectangle{
					X:      float32(i%l.mapData.Width) * TileSize,
					Y:      float32(i/l.mapData.Width) * TileSize,
					Width:  TileSize,
					Height: TileSize,
				}
				srcKey := GetAssetKey(tile.Tileset.Image.Source)
				srcImage, err := l.levelAssets.GetImage(TilesetLib, srcKey)
				if err != nil {
					return err
				}

				rl.DrawTexturePro(
					srcImage,
					srcRect,
					destRect,
					rl.Vector2{},
					0, rl.White,
				)
			}
		}
	}
	return nil
}
