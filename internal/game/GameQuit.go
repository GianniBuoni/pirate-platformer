package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *GameData) Quit() {
	g.unloadAssets()
	rl.CloseWindow()
}

func (g *GameData) unloadAssets() {
	g.levelAssets.Unload()
	// unload uiAssets
	// unload soundAssets
}
