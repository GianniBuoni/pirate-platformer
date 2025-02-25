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
	builder: sprites.NewSprite,
	groups:  []string{"all"},
}

var objectCollide = Loader[Object]{
	key:     "collision",
	builder: sprites.NewSprite,
	groups:  []string{"all", "collision"},
}

var playerLoader = Loader[Object]{
	key:     "player",
	builder: sprites.NewPlayer,
	groups:  []string{"player", "all"},
}

var platformLoader = Loader[Object]{
	key:     "platform",
	builder: sprites.NewMovingSprite,
	groups:  []string{"all", "moving", "platform"},
}

var animatedLoader = Loader[Object]{
	key:     "animated",
	builder: sprites.NewAnimatedSprite,
	groups:  []string{"all"},
}

var itemLoader = Loader[Object]{
	key:     "item",
	builder: sprites.NewAnimatedSprite,
	groups:  []string{"all", "item"},
}

var pathLoader = Loader[Object]{
	key:     "path",
	builder: sprites.NewPath,
	groups:  []string{"path"},
}

var damageLoader = Loader[Object]{
	key:     "damage",
	builder: sprites.NewMovingSprite,
	groups:  []string{"all", "moving", "damage"},
}
