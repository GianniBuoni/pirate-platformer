package game

import (
	"fmt"
	"os"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	"github.com/GianniBuoni/pirate-platformer/internal/level"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Load() {
	g.window = NewWindow()
	g.loadAssets()
	var err error
	mapPath := g.levelMaps[g.stats.currentLevel]
	g.levelCurrent, err = level.NewLevel(g.levelAssets, mapPath)
	if err != nil {
		fmt.Printf("❌: Game.Load(), could not init level %s\n", err.Error())
		os.Exit(2)
	}
	err = g.levelCurrent.Load(g.loaders)
	if err != nil {
		fmt.Printf("❌: Game.Load(), could not load level %s\n", err.Error())
		os.Exit(2)
	}
	g.window.loadCam(rl.NewVector2(lib.WindowW/2, lib.WindowH/2))
}

func (g *GameData) loadAssets() {
	// init assets
	g.levelAssets = assets.NewAssets()

	// load all assets
	assetMap := map[string]assets.AssetLibrary{
		"tilesets": assets.TilesetLib,
		"objects":  assets.ImageLib,
		"player":   assets.PlayerLib,
	}

	for k, v := range assetMap {
		err := g.levelAssets.ImportImages(v, "graphics", k)
		if err != nil {
			fmt.Printf("❌: Game.loadAssets(), %s: %s\n", k, err.Error())
			os.Exit(2)
		}
	}
	err := g.levelAssets.ImportTilesetData("data", "tilesets")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), tileset data: %s\n", err.Error())
		os.Exit(2)
	}
}
