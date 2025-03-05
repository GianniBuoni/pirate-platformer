package loaders

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type LoadersForLevel struct {
	Tiles   map[string]Loader[TileParams]
	Objects map[string]Loader[Object]
}

func LevelLoaders() LoadersForLevel {
	ll := LoadersForLevel{
		Tiles:   map[string]Loader[TileParams]{},
		Objects: map[string]Loader[Object]{},
	}
	levelTileLoaders := []Loader[TileParams]{
		&bgTileLoader, &cTileLoader, &pTileLoader,
	}
	levelObjectLoaders := []Loader[Object]{
		&objectLoader,
	}
	for _, v := range levelTileLoaders {
		ll.Tiles[v.Key()] = v
	}
	for _, v := range levelObjectLoaders {
		ll.Objects[v.Key()] = v
	}
	return ll
}
