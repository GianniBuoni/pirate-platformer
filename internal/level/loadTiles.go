package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

				sprite, err := sprites.NewSprite(
					srcKey, tilePos, l.levelAssets,
					sprites.WithImgPos(srcPos),
					sprites.WithAssetLib(TilesetLib),
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
