package game

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Run() {
	g.update()
	g.draw()
}

func (g *GameData) update() {
	g.running = !rl.WindowShouldClose()
	g.window.Update()

	for _, sprite := range g.allSprites {
		sprite.Update()
	}
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.GetRenderTexture())
	rl.ClearBackground(rl.Pink)

	err := DrawMap(g)
	if err != nil {
		fmt.Printf("❌: Game.draw(), DrawMap, %s", err.Error())
		os.Exit(2)
	}

	for _, sprite := range g.allSprites {
		sprite.Draw(g.levelAssets)
	}

	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()

}
