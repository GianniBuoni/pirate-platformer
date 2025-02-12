package game

import (
	"fmt"
	"path/filepath"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

var mapData *tiled.Map

func LoadMap() (err error) {
	mapPath := filepath.Join("data", "levels", "omni.tmx")
	mapData, err = tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error loading map: %s, %s", mapPath, err.Error())
	}
	fmt.Printf("W: %d, H: %d\n", mapData.Width, mapData.Height)
	return nil
}

func DrawMap(g *GameData) error {
	bgLayer := mapData.Layers[0]
	srcImage, err := g.levelAssets.GetImage("inside")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for i, tile := range bgLayer.Tiles {
		if !tile.IsNil() {
			srcData := tile.GetTileRect()
			srcRect := rl.Rectangle{
				X:      float32(srcData.Min.X),
				Y:      float32(srcData.Min.Y),
				Width:  lib.TileSize,
				Height: lib.TileSize}
			destRect := rl.Rectangle{
				X:      float32(i%mapData.Width) * lib.TileSize,
				Y:      float32(i/mapData.Width) * lib.TileSize,
				Width:  lib.TileSize,
				Height: lib.TileSize,
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
	return nil
}
