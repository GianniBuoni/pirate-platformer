package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameData) Run() {
  g.update()
  g.draw()
}

func (g *GameData) update() {
  g.running = !rl.WindowShouldClose()
  g.window.Update()
}

func (g *GameData) draw() {
  // darw onto render texture
  rl.BeginTextureMode(g.window.GetRenderTexture())
  rl.ClearBackground(rl.Pink)
  rl.EndTextureMode()

  // draw render texture scaled
  g.window.Draw()

}
