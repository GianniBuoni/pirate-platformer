package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) Draw() error {
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
	err := ui.drawStats()
	if err != nil {
		return err
	}
	return nil
}

func (ui *UI) drawStats() error {
	groups := []string{"coin", "heart", "ephemeral"}
	for _, name := range groups {
		if name == "heart" {
			rendered, err := ui.groups.GetIDs(name)
			if err != nil {
				return err
			}
			rendered = rendered[1:]
			hearts, err := ui.groups.GetSpritesID(name, rendered)
			if err != nil {
				return err
			}
			for _, s := range hearts {
				s.Draw(s.GetID().Src, s.GetPos())
			}
			continue
		}
		err := ui.groups.Draw(name)
		if err != nil {
			return err
		}
	}
	coins := fmt.Sprint(ui.stats.Coins)
	ui.texts["coinText"].Draw(coins, ui.assets.Fonts["runescape_uf"])
	return nil
}
