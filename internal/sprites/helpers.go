package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// call this func in the draw method to debug any hitbox and rect issues
func drawRect(r *Rect, c rl.Color) {
	rl.DrawRectangleRec(
		rl.Rectangle(*r), rl.ColorAlpha(c, .5),
	)
}
