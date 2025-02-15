package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) input() {
	var x float32
	var direction float32
	collision := p.checkCollisonSide()

	leftInput := []int32{rl.KeyLeft, rl.KeyH, rl.KeyA}
	rightInput := []int32{rl.KeyRight, rl.KeyL, rl.KeyD}

	// X direction is inverted by wall jumping
	// the Player.wallJump() method toggles this boolean
	if p.actions["run"] {
		direction = 1
	} else {
		direction = -1
	}

	for _, key := range leftInput {
		if rl.IsKeyDown(key) {
			x -= direction
		}
	}
	for _, key := range rightInput {
		if rl.IsKeyDown(key) {
			x += direction
		}
	}
	p.direction.X = x

	// Player.wallJump() params manually set since p.direction.X has
	// a chance of being 0.
	if rl.IsKeyPressed(rl.KeySpace) {
		switch collision {
		case floor:
			p.jump()
		case left:
			p.wallJump(1)
		case right:
			p.wallJump(-1)
		}
	}

	if rl.IsKeyPressed(rl.KeyF) {
		p.attack(collision)
	}
}
