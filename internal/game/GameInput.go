package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *GameData) input() {
	rl.SetExitKey(0)
	if rl.IsKeyPressed(rl.KeyEscape) {
		g.stats.Paused = !g.stats.Paused
	}

	if g.stats.Victory && rl.IsKeyPressed(rl.KeyEnter) {
		g.stats.EnterStage()
		g.LoadLevel()
	}
}
