package loaders

import (
	"math"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var bgTileLoader = SpriteLoader[MapProps]{
	Builder: bgMiddleware(NewTileSprite),
	Groups:  []string{"all"},
}

func bgMiddleware(
	f func(Tile, *Assets) (Sprite, error),
) func(MapProps, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(mp MapProps, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		if mp.Bg == "" {
			return nil, nil
		}
		extraTiles := int(math.Abs(float64(mp.TopLimit * mp.Width)))
		mapTiles := mp.Width * mp.Height
		totalTiles := mapTiles + extraTiles

		out := []Sprite{}
		for i := range totalTiles {
			x := float32(i%mp.Width) * TileSize
			y := float32(mp.TopLimit)*TileSize + float32(i/mp.Width)*TileSize
			bgTile := Tile{
				Image: mp.Bg,
				X:     x,
				Y:     y,
			}
			s, err := f(bgTile, gm.Assets())
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
