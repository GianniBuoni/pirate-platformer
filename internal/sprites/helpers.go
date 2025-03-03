package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// call this func in the draw method to debug any hitbox and rect issues
func drawRect(r *Rect, c rl.Color) {
	rl.DrawRectangleRec(
		rl.Rectangle(*r), rl.ColorAlpha(c, .5),
	)
}

// call printRect if drawRect isn't available
func printRect(r *Rect, pos string) {
	switch pos {
	case "top":
		fmt.Printf("Top: %v\n", r.Top())
	case "bottom":
		fmt.Printf("Bottom: %v\n", r.Bottom())
	default:
		fmt.Printf(
			`Top: %v, Bottom: %v
Left: %v, Right: %v
`,
			r.Top(), r.Bottom(), r.Left(), r.Right(),
		)
	}
}
