package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var cloudBGLoader = SpriteLoader[MapProps]{
	Builder: skyboxMiddleware(NewRectSrite, NewMovingSprite),
	Groups:  []string{"all", "moving", "clouds large"},
}

func skyboxMiddleware(
	f func(Object, rl.Color, *Assets) (Sprite, error),
	g func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(MapProps, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(mp MapProps, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		// check if valid map props
		if mp.Bg != "" {
			return nil, nil
		}

		// PROCESS MAP DATA HERE
		out := []Sprite{}

		// water rect
		o := Object{}
		o.Y = mp.Horizon * TileSize
		o.Width = float32(mp.Width) * TileSize
		o.Height = (float32(mp.Height) - mp.Horizon) * TileSize

		// call builder function
		s, err := f(o, WaterColor, gm.Assets())
		if err != nil {
			return nil, err
		}
		out = append(out, s)

		// sky
		for i := range 3 {
			cloud, err := gm.Assets().GetObject("large_cloud")
			if err != nil {
				return nil, err
			}
			s, err = g(cloud, ImageLib, gm.Assets())
			if err != nil {
				return nil, err
			}
			s.Rect().Set(Bottom(mp.Horizon * TileSize))
			s.Rect().Set(Left(float32(i) * s.Rect().Width))
			out = append(out, s)
		}
		return out, nil
	}
}
