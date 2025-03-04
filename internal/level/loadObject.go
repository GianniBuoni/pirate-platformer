package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoaders = []Loader[Object]{
	playerLoader, cloudLoader, pathLoader, sawPathLoader,
	platformLoader, damageLoader, radialLoader, itemLoader,
	shellLoader, waterLoader,
}

var itemLoader = Loader[Object]{
	key:     "item",
	builder: objectMiddleWare(NewItem),
	groups:  []string{"all", "item", "moving"},
}

var cloudLoader = Loader[Object]{
	key:     "cloud",
	builder: objectMiddleWare(NewCloud),
	groups:  []string{"all", "moving"},
}

func objectMiddleWare(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ImageLib, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		return []Sprite{s}, nil
	}
}
