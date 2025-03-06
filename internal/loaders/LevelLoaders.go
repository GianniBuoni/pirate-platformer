package loaders

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type LoadersForLevel struct {
	MapProps []Loader[MapProps]
	Tiles    map[string]Loader[TileParams]
	Objects  map[string]Loader[Object]
}

func LevelLoaders() LoadersForLevel {
	ll := LoadersForLevel{
		MapProps: []Loader[MapProps]{
			&bgTileLoader, &cloudBGLoader,
		},
		Tiles:   map[string]Loader[TileParams]{},
		Objects: map[string]Loader[Object]{},
	}
	levelTileLoaders := []Loader[TileParams]{
		&tileLoader, &cTileLoader, &pTileLoader,
	}
	levelObjectLoaders := []Loader[Object]{
		&objectLoader, &objectCollide, &animatedLoader, &radialLoader, &itemLoader,
		&waterLoader, &cloudLoader,
	}
	for _, v := range levelTileLoaders {
		ll.Tiles[v.GetKey()] = v
	}
	for _, v := range levelObjectLoaders {
		ll.Objects[v.GetKey()] = v
	}
	return ll
}
