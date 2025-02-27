package game

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Run() {
	g.update()
	g.draw()
}

func (g *GameData) update() {
	g.Running = !rl.WindowShouldClose()
	err := g.levelCurrent.Update()
	if err != nil {
		g.Quit(1, err)
	}
	g.window.Update(g.levelCurrent.PlayerPos())
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.renderTexture)
	rl.BeginMode2D(g.window.camera)
	rl.ClearBackground(BgColor)
	err := g.levelCurrent.Draw()
	if err != nil {
		g.Quit(1, err)
	}
	rl.EndMode2D()
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()
}
