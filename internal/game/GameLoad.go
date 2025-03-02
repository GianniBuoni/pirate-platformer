package game

import (
	"path/filepath"

	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/ui"
)

// LoadLevel is resposible for loading in levels
func (g *GameData) LoadLevel() {
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
}

func (g *GameData) LoadUi() {
	ui, err := ui.NewUI(g.stats, g.levelAssets)
	if err != nil {
		g.Quit(1, err)
	}
	g.ui = ui
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
	}

	for k, v := range graphics {
		err := g.levelAssets.ImportImages(v, "graphics", k)
		if err != nil {
			g.Quit(1, err)
		}
	}
	for k, v := range data {
		err := g.levelAssets.ImportData(v, "data", k)
		if err != nil {
			g.Quit(1, err)
		}
	}
}
