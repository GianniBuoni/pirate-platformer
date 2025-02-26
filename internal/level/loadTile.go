package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var tileLoaders = []Loader[Tile]{
	bgTileLoader, cTileLoader, pTileLoader,
}

var bgTileLoader = Loader[Tile]{
	key:     "bg",
	builder: tileMiddleware(sprites.NewTileSprite),
	groups:  []string{"all"},
}

var cTileLoader = Loader[Tile]{
	key:     "collision",
	builder: tileMiddleware(sprites.NewTileSprite),
	groups:  []string{"all", "collision", "wall"},
}

var pTileLoader = Loader[Tile]{
	key:     "platform",
	builder: tileMiddleware(sprites.NewTileSprite),
	groups:  []string{"all", "platform"},
}

func tileMiddleware(
	f func(Tile, *Assets) (Sprite, error),
) func(Tile, *LevelData) ([]Sprite, error) {
	return func(t Tile, ld *LevelData) ([]Sprite, error) {
		s, err := f(t, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		return []Sprite{s}, nil
	}
}
