package level

import (
	"errors"
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var pearlLoader = SpriteLoader[Object]{
	Key:     "pearl",
	Builder: pearlMiddleware(NewPearl),
	Groups:  []string{"damage", "ephemeral"},
}

func (l *Level) spawnPearl(s Sprite) error {
	obj, err := l.assets.GetObject("pearl")
	if err != nil {
		return err
	}
	switch s.GetPos().FlipH {
	case -1:
		obj.X = s.HitBox().Left()
		obj.Properties.DirX = -1
	default:
		obj.X = s.HitBox().Right()
	}
	obj.Y = s.HitBox().Center().Y
	pearlLoader.Run(obj, ImageLib, l)
	return nil
}

func pearlMiddleware(
	f func(Object, AssetLibrary, *Assets) (*Pearl, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		p, err := f(o, al, gm.Assets())
		if err != nil {
			return nil, err
		}
		level, ok := gm.(*Level)
		if !ok {
			return nil, errors.New("type mismatch, local Level loader used out of scope.")
		}
		p.Groups = level.groups
		p.Sprites = level.spirtes
		go func() {
			killTimer := time.NewTimer(p.Lifetime)
			<-killTimer.C
			p.GetID().Kill = true
		}()
		return []Sprite{p}, err
	}
}
