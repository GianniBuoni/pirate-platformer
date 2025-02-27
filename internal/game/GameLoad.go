package game

import (
	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (g *GameData) Load() {
	g.window = NewWindow()
	g.loadAssets()
	var err error
	mapPath := g.levelMaps[g.stats.currentLevel]
	g.levelCurrent, err = level.NewLevel(g.levelAssets, mapPath)
	if err != nil {
		g.Quit(1, err)
	}
	err = g.levelCurrent.Load(g.loaders)
	if err != nil {
		g.Quit(1, err)
	}
	g.window.loadCam(g.levelCurrent.PlayerPos())
}

func (g *GameData) loadAssets() {
	// init assets
	g.levelAssets = NewAssets()

	// load all assets
	assetMap := map[string]AssetLibrary{
		"tilesets": TilesetLib,
		"objects":  ImageLib,
		"player":   PlayerLib,
	}

	for k, v := range assetMap {
		err := g.levelAssets.ImportImages(v, "graphics", k)
		if err != nil {
			g.Quit(1, err)
		}
	}
	err := g.levelAssets.ImportTilesetData("data", "tilesets")
	if err != nil {
		g.Quit(1, err)
	}
	err = g.levelAssets.ImportSpawnIn("data", "templates")
	if err != nil {
		g.Quit(1, err)
	}
}
