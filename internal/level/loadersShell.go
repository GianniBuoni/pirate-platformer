package level

import (
	"errors"
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var shellLoader = SpriteLoader[Object]{
	Key:     "shell",
	Builder: shellMiddlware(NewShell),
	Groups:  []string{"all", "moving", "collision", "shell"},
}

func shellMiddlware(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		s, err := f(o, al, gm.Assets())
		if err != nil {
			return nil, err
		}
		shell, ok := s.(*Shell)
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
		shell.Player = level.player
		return []Sprite{shell}, nil
	}
}
