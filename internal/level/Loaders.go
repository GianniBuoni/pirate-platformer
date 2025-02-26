package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Loaders struct {
	tiles  map[string]Loader[Tile]
	object map[string]Loader[Object]
}

type Loader[T any] struct {
	key     string
	builder func(T, *LevelData) ([]Sprite, error)
	groups  []string
}

func (l *Loader[T]) Run(t T, level *LevelData) error {
	sprites, err := l.builder(t, level)
	if err != nil {
		return err
	}
	if sprites != nil {
		for _, s := range sprites {
			level.AddSpriteGroup(s, l.groups...)
		}
	}
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
