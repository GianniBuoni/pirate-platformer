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

var playerLoader = Loader[Object]{
	key:     "player",
	builder: objectMiddleWare(sprites.NewPlayer),
	groups:  []string{"player", "all"},
}

var platformLoader = Loader[Object]{
	key:     "platform",
	builder: objectMiddleWare(sprites.NewMovingSprite),
	groups:  []string{"all", "moving", "platform"},
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

var pathLoader = Loader[Object]{
	key:     "path",
	builder: objectMiddleWare(sprites.NewPath),
	groups:  []string{"path"},
}

var damageLoader = Loader[Object]{
	key:     "damage",
	builder: objectMiddleWare(sprites.NewMovingSprite),
	groups:  []string{"all", "moving", "damage"},
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
