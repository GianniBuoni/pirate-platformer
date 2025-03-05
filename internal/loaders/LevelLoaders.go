package loaders

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type LoadersForLevel struct {
	Tiles map[string]Loader[TileParams]
}

func LevelLoaders() LoadersForLevel {
	ll := LoadersForLevel{
		Tiles: map[string]Loader[TileParams]{},
	}
	levelTileLoaders := []Loader[TileParams]{
		&bgTileLoader, &cTileLoader, &pTileLoader,
	}
	for _, v := range levelTileLoaders {
		ll.Tiles[v.Key()] = v
	}
	return ll
}
