package level

import (
	"errors"
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var pathLoader = SpriteLoader[Object]{
	Key:     "path",
	Builder: pathMiddleware(NewPath),
}

var damageLoader = SpriteLoader[Object]{
	Key:     "damage",
	Builder: movingMiddleware(NewMovingSprite),
	Groups:  []string{"all", "moving", "damage"},
}

var platformLoader = SpriteLoader[Object]{
	Key:     "platform",
	Builder: movingMiddleware(NewMovingSprite),
	Groups:  []string{"all", "moving", "platform"},
}

func pathMiddleware(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
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
		return nil, nil
	}
}

func movingMiddleware(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		s, err := f(o, al, gm.Assets())
		if err != nil {
			return nil, err
		}
		moving, ok := s.(*MovingSprite)
		if !ok {
			return []Sprite{s}, nil
		}
		level, ok := gm.(*Level)
		if ok {
			moving.GetPath(level.paths)
		}
		return []Sprite{moving}, nil
	}
}
