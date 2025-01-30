package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *GameData) Run() {
  g.running = !rl.WindowShouldClose()
	rl.BeginDrawing()
	rl.ClearBackground(rl.Pink)
  rl.EndDrawing()
}
