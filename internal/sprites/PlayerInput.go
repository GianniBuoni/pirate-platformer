package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

func (p *Player) input(side CollisionSide) {
	var x float32

	leftInput := []int32{rl.KeyLeft, rl.KeyH, rl.KeyA}
	rightInput := []int32{rl.KeyRight, rl.KeyL, rl.KeyD}
	downInput := []int32{rl.KeyDown, rl.KeyJ, rl.KeyS}

	for _, key := range leftInput {
		if rl.IsKeyDown(key) {
			x -= 1
		}
	}
	for _, key := range rightInput {
		if rl.IsKeyDown(key) {
			x += 1
		}
	}
	if p.actions[run] {
		p.direction.X = x
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		switch side {
		case floor:
			p.jump()
		case left, right:
			p.wallJump(x * -1)
		}
	}

	for _, key := range downInput {
		if rl.IsKeyPressed(key) {
			p.phasePlatform()
		}
	}

	if rl.IsKeyPressed(rl.KeyF) {
		p.attack()
	}
}
