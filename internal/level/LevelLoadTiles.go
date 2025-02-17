package level

import (
	"strings"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadTiles(objs []*tiled.Object) error {
	// bg palms are zero index
	for _, obj := range objs {
		if strings.Contains(obj.Name, "palm") {
			s, err := sprites.NewSprite(
				obj.Name,
				rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets,
			)
			if err != nil {
				return err
			}
			l.AddSpriteGroup(s, "all")
		}
	}
	// all tile layers loaded together
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

				switch layer.Name {
				case "BG", "FG":
					l.AddSpriteGroup(sprite, "all")
				case "Platforms":
					l.AddSpriteGroup(sprite, "all", "platform")
				default:
					l.AddSpriteGroup(sprite, "all", "collision")
				}
			}
		}
	}
	return nil
}
