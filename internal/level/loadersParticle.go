package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/loaders"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *Level) spawnParticle(s sprites.Sprite) error {
	o, err := l.assets.GetObject("particle")
	if err != nil {
		return err
	}
	o.X = s.HitBox().Center().X
	o.Y = s.HitBox().Center().Y
	err = loaders.ParticleLoader.Run(o, lib.ImageLib, l)
	if err != nil {
		return err
	}
	return nil
}
