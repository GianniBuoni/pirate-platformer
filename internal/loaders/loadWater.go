package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var waterLoader = SpriteLoader[Object]{
	Key:     "water",
	Builder: waterMiddleware(NewAnimatedSprite, NewSprite),
	Groups:  []string{"all", "moving"},
}

func waterMiddleware(
	f, g func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
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
				s, err = f(newObject, ImageLib, gm.Assets())
			} else {
				s, err = g(newObject, ImageLib, gm.Assets())
			}
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		}
		return out, nil
	}
}
