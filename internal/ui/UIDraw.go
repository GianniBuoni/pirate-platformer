package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) Draw() error {
	// bg
	center := "GAME PAUSED"
	hint := "Press (ESC) again to resume."
	if ui.stats.PlayerHP() == 0 {
		center = "GAME OVER, MAN."
		hint = "Game over."
	}
	if ui.stats.Paused {
		rl.DrawRectangleRec(
			rl.NewRectangle(0, 0, WindowW, WindowH),
			rl.ColorAlpha(rl.Black, 0.7),
		)
		ui.texts["center message"].Draw(center, ui.assets.Fonts["runescape_uf"])
		ui.texts["center hint"].Draw(hint, ui.assets.Fonts["runescape_uf"])
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
