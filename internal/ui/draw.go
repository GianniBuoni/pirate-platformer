package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) Draw() {
	// bg
	rl.DrawRectangleRec(
		rl.NewRectangle(0, 0, WindowW, WindowH),
		rl.ColorAlpha(rl.Black, 0.3),
	)
	// stats
	for _, id := range ui.groups["all"] {
		s := ui.sprites[id]
		s.Draw(s.GetID().Src, s.GetPos())
	}
	ui.drawCoins()
}

func (ui *UI) drawCoins() {
	coins := fmt.Sprint(ui.stats.Coins)
	ui.texts["coinText"].Draw(coins)
}
