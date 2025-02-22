package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoader = Loader[Object]{
	Run: loadObject,
}

func loadObject(obj Object, l *LevelData) error {
	s, err := sprites.NewSprite(obj, l.levelAssets)
	if err != nil {
		return err
	}
	l.AddSpriteGroup(s, "all")
	return nil
}
