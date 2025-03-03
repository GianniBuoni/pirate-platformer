package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) drawStats() {
	// draw coins
	c, ok := ui.sprites["coin"]
	if ok {
		c.Draw(c.GetID().Src, c.GetPos())
	} else {
		fmt.Println("ui coin sprite not found")
	}

	coins := fmt.Sprintf("%d", ui.stats.Coins)
	ui.drawText(coins, rl.NewVector2(72, 64))

	// draw hearts
}

func loadHearts(ui *UI, dest []Sprite) {
	s := ui.sprites["heart"]
	pos := s.GetPos()
	for i := range ui.stats.PlayerHP() {
		var margin float32
		if i == 0 {
			margin = 32
		} else {
			margin = 32 + (float32(i) * 8)
		}
		pos.Rect().X = float32(i)*pos.Rect().Width + margin
	}

}

func (ui *UI) drawText(msg string, pos rl.Vector2) {
	rl.DrawTextEx(
		ui.assets.Fonts["runescape_uf"],
		msg, pos,
		32, 8, rl.White,
	)
}

func (ui *UI) drawTextDisplay(msg string, pos rl.Vector2) {
	rl.DrawTextEx(
		ui.assets.Fonts["runescape_uf"],
		msg, pos,
		64, 8, rl.White,
	)
}
