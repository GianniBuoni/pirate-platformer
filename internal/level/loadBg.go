package level

import (
	"math"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var propsLoaders = []Loader[MapProps]{
	loadBg,
}

var loadBg = Loader[MapProps]{
	builder: bgMiddleware(NewTileSprite),
	groups:  []string{"all"},
}

func bgMiddleware(
	f func(Tile, *Assets) (Sprite, error),
) func(MapProps, *LevelData) ([]Sprite, error) {
	return func(mp MapProps, ld *LevelData) ([]Sprite, error) {
		if mp.Bg == "" {
			return nil, nil
		}
		extraTiles := int(math.Abs(float64(mp.TopLimit * ld.Width)))
		mapTiles := ld.Height * ld.Width
		totalTiles := mapTiles + extraTiles

		out := []Sprite{}
		for i := range totalTiles {
			x := float32(i%ld.Width) * TileSize
			y := float32(mp.TopLimit)*TileSize + float32(i/ld.Width)*TileSize
			bgTile := Tile{
				Image: mp.Bg,
				X:     x,
				Y:     y,
			}
			s, err := f(bgTile, ld.levelAssets)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
