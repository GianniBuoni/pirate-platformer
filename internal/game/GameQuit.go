package game

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Quit(i int, errs ...error) {
	g.unloadAssets()
	rl.CloseWindow()
	for _, err := range errs {
		fmt.Printf("❌ ERROR: %s\n", err.Error())
	}
	os.Exit(i)
}

func (g *GameData) unloadAssets() {
	g.levelAssets.Unload()
	// unload uiAssets
	// unload soundAssets
}
