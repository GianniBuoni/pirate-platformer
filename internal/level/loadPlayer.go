package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var playerLoader = Loader[Object]{
	key:     "player",
	builder: playerMiddleware(NewPlayer),
}

func playerMiddleware(
	f func(Object, *Stats, *Assets) (Sprite, error),
) func(Object, *LevelData) ([]Sprite, error) {
	return func(o Object, ld *LevelData) ([]Sprite, error) {
		s, err := f(o, ld.stats, ld.levelAssets)
		if err != nil {
			return nil, err
		}
		player, ok := s.(*Player)
		if !ok {
			return nil, fmt.Errorf(
				"error building player: %s, is a Player Sprite",
				o.Image,
			)
		}
		player.Groups = ld.groups
		ld.AddPlayer(player)
		return nil, nil
	}
}
