package game

import (
	"path/filepath"

	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/ui"
)

func (g *GameData) LoadLevel() {
	g.level = level.NewLevel(g.stats, g.assets)
	err := g.level.Load(g.assets.Maps[g.stats.CurrentLevel])
	if err != nil {
		g.Quit(1, err)
	}
}

func (g *GameData) LoadUi() {
	var err error
	g.ui, err = ui.NewUI(g.stats, g.assets)
	if err != nil {
		g.Quit(1, err)
	}
	mapPath := filepath.Join("data", "ui", "base.json")
	err = g.ui.Load(mapPath)
	if err != nil {
		g.Quit(1, err)
	}
}

func (g *GameData) loadAssets() {
	// load all assets
	graphics := map[string]AssetLibrary{
		"tilesets": TilesetLib,
		"level":    ImageLib,
		"player":   PlayerLib,
		"ui":       UiLib,
	}

	data := map[string]AssetLibrary{
		"tilesets":  TileData,
		"templates": SpawnInLib,
		"fonts":     FontLib,
		"levels":    MapLib,
	}
	for k, v := range data {
		err := g.assets.ImportData(v, "data", k)
		if err != nil {
			g.Quit(1, err)
		}
	}
	for k, v := range graphics {
		err := g.assets.ImportImages(v, "graphics", k)
		if err != nil {
			g.Quit(1, err)
		}
	}
}
