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
	g.input()
	err := g.level.Update()
	if err != nil {
		g.Quit(1, err)
	}
	g.ui.Update()
	g.window.Update(g.level.CameraPos())
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.renderTexture)
	rl.BeginMode2D(g.window.camera)
	rl.ClearBackground(BgColor)
	g.level.Draw()
	rl.EndMode2D()
	err := g.ui.Draw()
	if err != nil {
		g.Quit(1, err)
	}
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()
}
