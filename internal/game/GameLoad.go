package game

import (
	"fmt"
	"os"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	"github.com/GianniBuoni/pirate-platformer/internal/window"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Load() {
	g.window = window.NewWindow()
	g.loadAssets()

	// start replace with level load
	boat, err := sprites.NewSprite(
		"boat",
		rl.NewVector2(lib.WindowW/2, lib.WindowH/2),
		g.levelAssets,
	)
	if err != nil {
		fmt.Println(err)
		// todo figure out how to trigger data unload on error
	}

	boat.OffsetCentre()
	g.AddSprite(boat)

	player, err := sprites.NewPlayer(rl.NewVector2(lib.WindowW/2, lib.WindowH/2-86), g.levelAssets)
	if err != nil {
		fmt.Println(err)
		// todo figure out how to trigger data unload on error
	}
	g.AddPlayer(player)

	// end replace with level load
}

func (g *GameData) loadAssets() {
	// init assets
	g.levelAssets = assets.NewAssets()

	// load all assets
	g.levelAssets.ImportImages("graphics", "objects")
	g.levelAssets.ImportPlayer("graphics", "player")
	g.levelAssets.ImportImages("graphics", "tilesets")

	err := LoadMap()
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}

}
