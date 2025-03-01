package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var loadSky = Loader[MapProps]{
	builder: skyboxMiddleware(NewMovingSprite),
	groups:  []string{"all", "moving"},
}

func skyboxMiddleware(
	f func(Object, *Assets) (Sprite, error),
) func(MapProps, *LevelData) ([]Sprite, error) {
	return func(mp MapProps, ld *LevelData) ([]Sprite, error) {
		// check if valid map props
		if mp.Bg == "" {
			return nil, nil
		}

		// PROCESS MAP DATA HERE
		out := []Sprite{}
		o := Object{}

		// call builder function
		s, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
		return nil, nil
	}
}
