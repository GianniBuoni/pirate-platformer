package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var particleLoader = Loader[Object]{
	key:     "particle",
	builder: objectMiddleWare(NewParticle),
	groups:  []string{"all", "moving"},
}

func (l *LevelData) spawnParticle(s Sprite) error {
	o, err := l.levelAssets.GetObject("particle")
	if err != nil {
		return err
	}
	o.X = s.HitBox().Center().X
	o.Y = s.HitBox().Center().Y
	err = particleLoader.Run(o, l)
	if err != nil {
		return err
	}
	return nil
}
