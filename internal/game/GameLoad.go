package game

import (
	"fmt"
	"os"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/level"
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
	//g.window.loadCam(g.levelCurrent.PlayerPos())
}

func (g *GameData) loadAssets() {
	// init assets
	g.levelAssets = assets.NewAssets()

	// load all assets
	err := g.levelAssets.ImportImages(ImageLib, "graphics", "objects")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), images: %s", err.Error())
		os.Exit(1)
	}
	err = g.levelAssets.ImportImages(PlayerLib, "graphics", "player")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), player: %s", err.Error())
		os.Exit(1)
	}
	err = g.levelAssets.ImportImages(TilesetLib, "graphics", "tilesets")
	if err != nil {
		fmt.Printf("❌: Game.loadAssets(), tilesets: %s", err.Error())
		os.Exit(1)
	}
}
