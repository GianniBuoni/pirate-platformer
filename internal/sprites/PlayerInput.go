package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

func (p *Player) input() {
	var x float32
	var dir float32
	dir = 1

	leftInput := []int32{rl.KeyLeft, rl.KeyH, rl.KeyA}
	rightInput := []int32{rl.KeyRight, rl.KeyL, rl.KeyD}

	for _, key := range leftInput {
		if rl.IsKeyDown(key) {
			x -= dir
		}
	}
	for _, key := range rightInput {
		if rl.IsKeyDown(key) {
			x += dir
		}
	}
	p.direction.X = x

	if rl.IsKeyPressed(rl.KeySpace) {
		p.jump()
	}
}
