package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
