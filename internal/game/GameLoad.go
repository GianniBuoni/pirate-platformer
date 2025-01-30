package game

import (
	"log"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	"github.com/GianniBuoni/pirate-platformer/internal/window"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Load() {
	g.window = window.NewWindow()
	g.levelAssets = assets.NewAssets()
	g.loadAssets()
	boat, err := sprites.NewSprite(
		"boat",
		rl.NewVector2(lib.WindowW/2, lib.WindowH/2),
		g.levelAssets,
	)
	if err != nil {
		log.Fatal(err)
	}

  boat.OffsetCentre()
	g.AddSprite(boat)

}

func (g *GameData) loadAssets() {
	g.levelAssets.ImportFolder([]string{"graphics", "objects"})
}
