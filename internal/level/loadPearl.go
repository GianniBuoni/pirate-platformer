package level

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var pearlLoader = Loader[Object]{
	key:     "pearl",
	builder: pearlMiddleware(NewPearl),
	groups:  []string{"all", "moving", "damage", "pearl"},
}

func (l *LevelData) spawnPearl(s Sprite) error {
	obj, err := l.levelAssets.GetObject("pearl")
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
	pearlLoader.Run(obj, l)
	return nil
}

func pearlMiddleware(
	f func(Object, *Assets) (*Pearl, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		p, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		p.Groups = ld.groups
		go func() {
			killTimer := time.NewTimer(p.Lifetime)
			<-killTimer.C
			p.GetID().Kill = true
		}()
		return []Sprite{p}, err
	}
}
