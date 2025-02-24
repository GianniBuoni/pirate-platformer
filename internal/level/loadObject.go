package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoaders = []Loader[Object]{
	animatedLoader,
	objectLoader,
	objectCollide,
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
	builder: sprites.NewAnimatedSprite,
	groups:  []string{"all", "moving", "platform"},
}

var animatedLoader = Loader[Object]{
	key:     "animated",
	builder: sprites.NewAnimatedSprite,
	groups:  []string{"all"},
}
