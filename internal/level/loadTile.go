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
	builder: sprites.NewTileSprite,
	groups:  []string{"all"},
}

var cTileLoader = Loader[Tile]{
	key:     "collision",
	builder: sprites.NewTileSprite,
	groups:  []string{"all", "collision", "wall"},
}

var pTileLoader = Loader[Tile]{
	key:     "platform",
	builder: sprites.NewTileSprite,
	groups:  []string{"all", "platform"},
}
