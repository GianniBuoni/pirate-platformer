package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) Draw() {
	// bg
	if ui.stats.Paused {
		rl.DrawRectangleRec(
			rl.NewRectangle(0, 0, WindowW, WindowH),
			rl.ColorAlpha(rl.Black, 0.3),
		)
		ui.texts["pause message"].Draw(
			"GAME PAUSED", ui.assets.Fonts["runescape_uf"],
		)
		ui.texts["pause hint"].Draw(
			"Press (ESC) again to resume.", ui.assets.Fonts["runescape_uf"],
		)
	}
	// stats
	ui.drawStats()
}

func (ui *UI) drawStats() {
	groups := []string{"coin", "heart", "ephemeral"}
	for _, name := range groups {
		group, ok := ui.groups[name]
		if !ok {
			continue
		}
		if name == "heart" {
			group = group[1:]
		}
		if len(group) == 0 {
			continue
		}
		for _, id := range group {
			s := ui.sprites[id]
			s.Draw(s.GetID().Src, s.GetPos())
		}
	}
	coins := fmt.Sprint(ui.stats.Coins)
	ui.texts["coinText"].Draw(coins, ui.assets.Fonts["runescape_uf"])
}
