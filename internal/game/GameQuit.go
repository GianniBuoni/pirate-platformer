package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *GameData) Quit() {
  rl.CloseWindow()
}
