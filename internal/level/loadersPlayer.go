package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var playerLoader = SpriteLoader[Object]{
	Key:     "player",
	Builder: playerMiddleware(NewPlayer),
}

func playerMiddleware(
	f func(Object, *Stats, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		ld, ok := gm.(*Level)
		s, err := f(o, ld.stats, ld.assets)
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
		ld.addPlayer(player)
		return nil, nil
	}
}
