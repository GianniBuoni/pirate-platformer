package level

import (
	"errors"
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var sawPathLoader = SpriteLoader[Object]{
	Key:     "sawPath",
	Builder: sawPathMiddleware(NewPath, NewSprite),
	Groups:  []string{"all"},
}

func sawPathMiddleware(
	f, g func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		// path
		s, err := f(o, al, gm.Assets())
		if err != nil {
			return nil, err
		}
		path, ok := s.(*ObjectSprite)
		if !ok {
			return nil, fmt.Errorf(
				"sprite \"%s\" is not path sprite. Wrong loader assigned in tiled\n",
				o.Image,
			)
		}
		level, ok := gm.(*Level)
		if !ok {
			return nil, errors.New(
				"type mismatch, non Level game module using level only loader\n",
			)
		}
		err = path.SetPaths(level.paths)
		if err != nil {
			return nil, err
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
		src, err := level.assets.GetImage(ImageLib, img)
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
			s, err := g(newO, al, level.assets)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
