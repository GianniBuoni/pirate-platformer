package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoaders = []Loader[Object]{
	animatedLoader,
	damageLoader,
	itemLoader,
	objectLoader,
	objectCollide,
	pathLoader,
	radialLoader,
	platformLoader,
	playerLoader,
	sawPathLoader,
	shellLoader,
}

var objectLoader = Loader[Object]{
	key:     "object",
	builder: objectMiddleWare(sprites.NewSprite),
	groups:  []string{"all"},
}

var objectCollide = Loader[Object]{
	key:     "collision",
	builder: objectMiddleWare(sprites.NewSprite),
	groups:  []string{"all", "collision"},
}

var animatedLoader = Loader[Object]{
	key:     "animated",
	builder: objectMiddleWare(sprites.NewAnimatedSprite),
	groups:  []string{"all"},
}

var itemLoader = Loader[Object]{
	key:     "item",
	builder: objectMiddleWare(sprites.NewAnimatedSprite),
	groups:  []string{"all", "item"},
}

func objectMiddleWare(
	f func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		return []Sprite{s}, nil
	}
}
