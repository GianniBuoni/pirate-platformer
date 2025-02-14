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
	g.levelCurrent.Update()
	g.window.Update(g.levelCurrent.PlayerPos())
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.renderTexture)
	rl.BeginMode2D(g.window.camera)
	rl.ClearBackground(rl.Pink)
	err := g.levelCurrent.Draw()
	if err != nil {
		fmt.Printf("Game.draw(), %s", err.Error())
		os.Exit(2)
	}
	rl.EndMode2D()
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()
}
