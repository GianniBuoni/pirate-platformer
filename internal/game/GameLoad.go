package game

import (
	"github.com/GianniBuoni/pirate-platformer/internal/window"
)

func (g *GameData) Load() {
  g.window = window.NewWindow()
}
