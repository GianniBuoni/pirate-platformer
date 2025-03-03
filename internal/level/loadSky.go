package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var loadSky = Loader[MapProps]{
	builder: skyboxMiddleware(NewRectSrite, NewMovingSprite),
	groups:  []string{"all", "moving", "clouds large"},
}

func skyboxMiddleware(
	f func(Object, rl.Color, *Assets) (Sprite, error),
	g func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(MapProps, *LevelData) ([]Sprite, error) {
	return func(mp MapProps, ld *LevelData) ([]Sprite, error) {
		// check if valid map props
		if mp.Bg != "" {
			return nil, nil
		}

		// PROCESS MAP DATA HERE
		out := []Sprite{}

		// water rect
		o := Object{}
		o.Y = mp.Horizon * TileSize
		o.Width = float32(ld.Width) * TileSize
		o.Height = (float32(ld.Height) - mp.Horizon) * TileSize

		// call builder function
		s, err := f(o, WaterColor, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		out = append(out, s)

		// sky
		for i := range 3 {
			cloud, err := ld.levelAssets.GetObject("large_cloud")
			if err != nil {
				return nil, err
			}
			s, err = g(cloud, ImageLib, ld.levelAssets)
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
