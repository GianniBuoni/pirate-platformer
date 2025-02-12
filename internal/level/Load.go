package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *LevelData) Load() error {
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
				l.AddSprite(sprite)
				// TODO add conditional to add sprite to collisionSprites
			}
		}
	}
	return nil
}
