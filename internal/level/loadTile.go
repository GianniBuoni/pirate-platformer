package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var tileLoaders = []Loader[[]int]{
	bgTileLoader, cTileLoader, pTileLoader,
}

var bgTileLoader = Loader[[]int]{
	key:     "bg",
	builder: tileMiddleware(NewTileSprite),
	groups:  []string{"all"},
}

var cTileLoader = Loader[[]int]{
	key:     "collision",
	builder: tileMiddleware(NewTileSprite),
	groups:  []string{"all", "collision", "wall"},
}

var pTileLoader = Loader[[]int]{
	key:     "platform",
	builder: tileMiddleware(NewTileSprite),
	groups:  []string{"all", "platform"},
}

func tileMiddleware(
	f func(Tile, *Assets) (Sprite, error),
) func([]int, *LevelData) ([]Sprite, error) {
	return func(data []int, ld *LevelData) ([]Sprite, error) {
		out := []Sprite{}
		for idx, id := range data {
			if id == 0 {
				continue
			}
			tile, err := parseTile(idx, id, ld)
			if err != nil {
				return nil, err
			}
			s, err := f(tile, ld.levelAssets)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
