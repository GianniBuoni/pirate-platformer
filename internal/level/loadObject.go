package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoaders = []Loader[Object]{
	playerLoader, cloudLoader, objectLoader, objectCollide, pathLoader, sawPathLoader,
	animatedLoader, platformLoader, damageLoader, radialLoader, itemLoader,
	shellLoader, waterLoader,
}

var objectLoader = Loader[Object]{
	key:     "object",
	builder: objectMiddleWare(NewSprite),
	groups:  []string{"all"},
}

var objectCollide = Loader[Object]{
	key:     "collision",
	builder: objectMiddleWare(NewSprite),
	groups:  []string{"all", "collision"},
}

var animatedLoader = Loader[Object]{
	key:     "animated",
	builder: objectMiddleWare(NewAnimatedSprite),
	groups:  []string{"all", "moving"},
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
