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
	g.levelCurrent, err = level.NewLevel(g.levelAssets, g.levelMaps[g.stats.currentLevel])
	if err != nil {
		fmt.Printf("❌: Game.Load(), could not init level %s", err.Error())
		os.Exit(2)
	}
	err = g.levelCurrent.Load()
	if err != nil {
		fmt.Printf("❌: Game.Load(), could not load level %s", err.Error())
		os.Exit(2)
	}
	g.window.loadCam(rl.NewVector2(lib.WindowW/2, lib.WindowH/2))
}

func (g *GameData) loadAssets() {
	// init assets
	g.levelAssets = assets.NewAssets()

	// load all assets
	err := g.levelAssets.ImportImages(assets.ImageLib, "graphics", "objects")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), images: %s", err.Error())
		os.Exit(1)
	}
	err = g.levelAssets.ImportImages(assets.PlayerLib, "graphics", "player")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), player: %s", err.Error())
		os.Exit(1)
	}
	err = g.levelAssets.ImportImages(assets.TilesetLib, "graphics", "tilesets")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), tilesets: %s", err.Error())
		os.Exit(1)
	}
	err = g.levelAssets.ImportTilesetData("data", "tilesets")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), tileset data: %s", err.Error())
		os.Exit(1)
	}
}
