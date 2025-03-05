package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var bgTileLoader = SpriteLoader[TileParams]{
	key:     "bg",
	Builder: tileMiddleware(NewTileSprite),
	Groups:  []string{"all"},
}

var cTileLoader = SpriteLoader[TileParams]{
	key:     "collision",
	Builder: tileMiddleware(NewTileSprite),
	Groups:  []string{"all", "collsion", "wall"},
}

var pTileLoader = SpriteLoader[TileParams]{
	key:     "platform",
	Builder: tileMiddleware(NewTileSprite),
	Groups:  []string{"all", "platform"},
}

func tileMiddleware(
	f func(Tile, *Assets) (Sprite, error),
) func(TileParams, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(tp TileParams, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		out := []Sprite{}
		for i := range tp.Data {
			if tp.Data[i] == 0 {
				continue
			}
			tile, err := ParseTileImage(tp, i, gm.Assets())
			if err != nil {
				return nil, err
			}
			s, err := f(tile, gm.Assets())
			if err != nil {
				return nil, err
			}
			s.GetID().GID = gm.NextId()
			out = append(out, s)
		}
		return out, nil
	}
}
