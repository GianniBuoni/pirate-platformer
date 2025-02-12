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
	g.Running = !rl.WindowShouldClose()
	g.window.Update()
	g.levelCurrent.Update()
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.GetRenderTexture())
	rl.ClearBackground(rl.Pink)
	err := g.levelCurrent.Draw()
	if err != nil {
		fmt.Printf("Game.draw(), %s", err.Error())
		os.Exit(2)
	}
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()

}
