package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var objectLoader = SpriteLoader[Object]{
	Key:     "object",
	Builder: objectMiddleWare(NewSprite),
	Groups:  []string{"all"},
}

var objectCollide = SpriteLoader[Object]{
	Key:     "collision",
	Builder: objectMiddleWare(NewSprite),
	Groups:  []string{"all", "collision"},
}

var HeartLoader = SpriteLoader[Object]{
	Key:     "heart",
	Builder: objectMiddleWare(NewHeartSprite),
	Groups:  []string{"heart"},
}

var coinLoader = SpriteLoader[Object]{
	Key:     "coin",
	Builder: objectMiddleWare(NewSprite),
	Groups:  []string{"coin"},
}

func objectMiddleWare(
	f func(Object, AssetLibrary, *Assets) (Sprite, error),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, a AssetLibrary, gm GameModule) ([]Sprite, error) {
		s, err := f(o, a, gm.Assets())
		if err != nil {
			return nil, err
		}
		return []Sprite{s}, nil
	}
}
