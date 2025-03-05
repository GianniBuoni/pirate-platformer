package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var radialLoader = SpriteLoader[Object]{
	Key:     "radial",
	Builder: spikeMiddleware(NewRadialSprite),
	Groups:  []string{"all", "moving"},
}

func spikeMiddleware(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		// spike ball sprite
		sprites := []Sprite{}
		spike, err := f(o, ImageLib, gm.Assets())
		if err != nil {
			return nil, err
		}
		gm.AddSpriteGroup(spike, gm.Sprites(), "damage")

		// chain sprites
		newImage := "spike_chain"
		src, err := gm.Assets().GetImage(ImageLib, newImage)
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
			s, err := f(newO, ImageLib, gm.Assets())
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
