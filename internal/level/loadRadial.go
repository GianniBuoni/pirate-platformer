package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var radialLoader = Loader[Object]{
	key:     "radial",
	builder: spikeMiddleware(sprites.NewRadialSprite),
	groups:  []string{"all", "moving"},
}

func spikeMiddleware(
	f func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		// spike ball sprite
		sprites := []Sprite{}
		spike, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		ld.AddSpriteGroup(spike, "damage")

		// chain sprites
		newImage := "spike_chain"
		src, err := ld.levelAssets.GetImage(ImageLib, newImage)
		if err != nil {
			return nil, err
		}
		center := spike.Rect().Center()
		maxR := o.Properties.DirX
		for i := float32(0); i < maxR; i += 20 {
			newO := o
			newO.Image = newImage
			newO.Width = float32(src.Width)
			newO.Height = float32(src.Height)
			newO.Properties.DirX = i
			s, err := f(newO, ld.levelAssets)
			if err != nil {
				return nil, err
			}
			s.OldRect().Set(Center(center.X, center.Y))
			sprites = append(sprites, s)
		}
		sprites = append(sprites, spike)

		return sprites, nil
	}
}
