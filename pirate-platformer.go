package main

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// init asset structs
	levelAssets := &lib.Assets{
		Frames: map[string][]rl.Texture2D{},
	}

	// init raylib
	rl.InitWindow(800, 450, "Super Pirate Platformer!")
	defer rl.CloseWindow()

	// load assets
	levelAssets.ImportImage([]string{"graphics", "objects", "boat"})
	levelAssets.ImportFolder([]string{"graphics", "items", "gold"})
	fmt.Println(levelAssets.Frames)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Pink)

		rl.EndDrawing()
	}
}
