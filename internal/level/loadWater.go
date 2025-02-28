package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var waterLoader = Loader[Object]{
	key:     "water",
	builder: waterMiddleware(NewAnimatedSprite, NewSprite),
	groups:  []string{"all"},
}

func waterMiddleware(
	f, g func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		width := o.Width / TileSize
		height := o.Height / TileSize
		tileCount := int(width * height)

		o.Width = TileSize
		o.Height = TileSize

		out := []Sprite{}
		for i := range tileCount {
			x := float32(i % int(width) * int(TileSize))
			y := float32(i / int(width) * int(TileSize))

			topRow := (i / int(width)) == 0
			newObject := o
			newObject.X += x
			newObject.Y += y

			var (
				s   Sprite
				err error
			)
			if topRow {
				newObject.Image = "water_surface"
				s, err = f(newObject, ld.levelAssets)
			} else {
				s, err = g(newObject, ld.levelAssets)
			}
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
