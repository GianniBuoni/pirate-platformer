package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var shellLoader = Loader[Object]{
	key:     "shell",
	builder: shellMiddlware(NewShell),
	groups:  []string{"all", "moving", "collision", "shell"},
}

func shellMiddlware(
	f func(Object, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		shell, ok := s.(*Shell)
		if ok {
			shell.Player = ld.player
		}
		return []Sprite{shell}, nil
	}
}
