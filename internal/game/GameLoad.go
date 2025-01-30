package game

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	"github.com/GianniBuoni/pirate-platformer/internal/window"
)

func (g *GameData) Load() {
	g.window = window.NewWindow()
	g.levelAssets = assets.NewAssets()
	g.loadAssets()
}

func (g *GameData) loadAssets() {
	g.levelAssets.ImportFolder([]string{"graphics", "objects"})
}
