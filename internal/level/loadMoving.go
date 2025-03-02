package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var pathLoader = Loader[Object]{
	key:     "path",
	builder: pathMiddleware(NewPath),
}

var damageLoader = Loader[Object]{
	key:     "damage",
	builder: movingMiddleware(NewMovingSprite),
	groups:  []string{"all", "moving", "damage"},
}

var platformLoader = Loader[Object]{
	key:     "platform",
	builder: movingMiddleware(NewMovingSprite),
	groups:  []string{"all", "moving", "platform"},
}

func pathMiddleware(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ImageLib, ld.levelAssets)
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
		return nil, nil
	}
}

func movingMiddleware(
	f func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		moving, ok := s.(*MovingSprite)
		if ok {
			moving.GetPath(ld.pathRects)
		}
		return []Sprite{moving}, nil
	}
}
