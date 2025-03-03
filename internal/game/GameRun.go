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
	g.ui.Update()
	g.window.Update(
		rl.NewVector2(WindowW/2, WindowH/2),
	)
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.renderTexture)
	rl.BeginMode2D(g.window.camera)
	rl.ClearBackground(BgColor)
	rl.EndMode2D()
	g.ui.Draw()
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()
}
