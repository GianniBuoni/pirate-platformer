package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var radialLoader = Loader[Object]{
	key:     "radial",
	builder: sprites.NewRadialSprite,
	groups:  []string{"all", "moving", "damage"},
}
