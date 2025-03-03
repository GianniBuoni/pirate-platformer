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
	g.input()
	g.Running = !rl.WindowShouldClose()
	if g.CanUpdate {
		err := g.levelCurrent.Update()
		if err != nil {
			g.Quit(1, err)
		}
		g.window.Update(g.levelCurrent.CameraPos())
		g.ui.Update()
	}
}

func (g *GameData) draw() {
	// draw onto render texture
	rl.BeginTextureMode(g.window.renderTexture)
	rl.BeginMode2D(g.window.camera)
	rl.ClearBackground(BgColor)
	g.levelCurrent.Draw()
	rl.EndMode2D()
	g.ui.Draw(g.CanUpdate)
	rl.EndTextureMode()

	// draw render texture scaled
	g.window.Draw()
}
