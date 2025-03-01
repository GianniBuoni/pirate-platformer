package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var loadSky = Loader[MapProps]{
	builder: skyboxMiddleware(NewRectSrite),
	groups:  []string{"all"},
}

func skyboxMiddleware(
	f func(Object, rl.Color, rl.Color, *Assets) (Sprite, error),
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
		s, err := f(o, WaterColor, rl.White, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
		return out, nil
	}
}
