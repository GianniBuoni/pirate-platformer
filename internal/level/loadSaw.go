package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var sawPathLoader = Loader[Object]{
	key:     "sawPath",
	builder: sawPathMiddleware(NewPath, NewSprite),
	groups:  []string{"all"},
}

func sawPathMiddleware(
	f func(Object, *Assets) (Sprite, error),
	g func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		// path
		s, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		path, ok := s.(*ObjectSprite)
		if ok {
			err = path.SetPaths(ld.pathRects)
			if err != nil {
				return nil, err
			}
		}

		// path sprites
		out := []Sprite{}
		var (
			max   float32
			moveH bool
		)

		if o.Width > o.Height {
			max = o.Width
			moveH = true
		} else {
			max = o.Height
			moveH = false
		}

		img := "saw_chain"
		src, err := ld.levelAssets.GetImage(ImageLib, img)
		if err != nil {
			return nil, err
		}
		o.Image = img
		o.Width = float32(src.Width)
		o.Height = float32(src.Height)

		for i := float32(0); i < max; i += 20 {
			newO := o
			if moveH {
				newO.X += i
			} else {
				newO.Y += i
			}
			s, err := g(newO, ld.levelAssets)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
