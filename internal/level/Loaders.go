package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Loaders struct {
	tiles  map[string]Loader[Tile]
	object map[string]Loader[Object]
}

type Loader[T any] struct {
	key     string
	builder func(T, *assets.Assets) (Sprite, error)
	groups  []string
}

func (l *Loader[T]) Run(t T, level *LevelData) error {
	s, err := l.builder(t, level.levelAssets)
	if err != nil {
		return err
	}
	level.AddSpriteGroup(s, l.groups...)
	return nil
}

func NewLoaders() *Loaders {
	l := Loaders{
		tiles:  map[string]Loader[Tile]{},
		object: map[string]Loader[Object]{},
	}
	for _, loader := range tileLoaders {
		l.tiles[loader.key] = loader
	}
	for _, loader := range objectLoaders {
		l.object[loader.key] = loader
	}
	return &l
}
