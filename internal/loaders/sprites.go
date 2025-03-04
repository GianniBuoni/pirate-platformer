package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var ObjectLoader = SpriteLoader[Object]{
	Key:     "object",
	Builder: ObjectMiddleWare(NewSprite),
	Groups:  []string{"all"},
}

func ObjectMiddleWare(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, a AssetLibrary, gm GameModule) ([]Sprite, error) {
		s, err := f(o, a, gm.Assets())
		if err != nil {
			return nil, err
		}
		s.GetID().GID = gm.NextId()
		return []Sprite{s}, nil
	}
}
