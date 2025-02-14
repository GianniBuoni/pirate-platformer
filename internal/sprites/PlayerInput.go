package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) input() {
	var x float32
	left := []int32{rl.KeyLeft, rl.KeyH, rl.KeyA}
	right := []int32{rl.KeyRight, rl.KeyL, rl.KeyD}
	for _, key := range left {
		if rl.IsKeyDown(key) {
			x -= 1
		}
	}
	for _, key := range right {
		if rl.IsKeyDown(key) {
			x += 1
		}
	}
	p.direction.X = x

	if rl.IsKeyPressed(rl.KeySpace) {
		p.jump()
	}
}
