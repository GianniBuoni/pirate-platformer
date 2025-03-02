package game

import (
	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/ui"
)

func (g *GameData) Load() {
	g.window = NewWindow()
	g.loadAssets()
	var err error
	mapPath := g.levelMaps[g.stats.CurrentLevel]
	g.levelCurrent, err = level.NewLevel(g.stats, g.levelAssets, mapPath)
	if err != nil {
		g.Quit(1, err)
	}
	err = g.levelCurrent.Load(g.loaders)
	if err != nil {
		g.Quit(1, err)
	}
	g.window.loadCam(g.levelCurrent.CameraPos())
	g.ui = ui.NewUI(g.stats, g.levelAssets)
}

func (g *GameData) loadAssets() {
	// load all assets
	assetMap := map[string]AssetLibrary{
		"tilesets": TilesetLib,
		"level":    ImageLib,
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
