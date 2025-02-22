package level

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Loaders struct {
	tiles  map[string]Loader[[]int]
	object map[string]Loader[Object]
}
type Loader[T any] struct {
	Key string
	Run func(T, *LevelData) error
}

func NewLoaders() *Loaders {
	// registered loaders
	tiles := []Loader[[]int]{bgTileLoader, cTileLoader, pTileLoader}
	object := []Loader[Object]{objectLoader}

	l := Loaders{
		tiles:  map[string]Loader[[]int]{},
		object: map[string]Loader[Object]{},
	}
	for _, loader := range tiles {
		l.tiles[loader.Key] = loader
	}
	for _, loader := range object {
		l.object[loader.Key] = loader
	}
	return &l
}
