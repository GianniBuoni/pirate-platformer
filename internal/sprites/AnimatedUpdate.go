package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

func (as *AnimatedSprite) Update() {
	dt := rl.GetFrameTime()
	as.rect.X += as.direction.X * as.speed * dt
	as.rect.Y += as.direction.Y * as.speed * dt
	as.constrain()
}

func (as *AnimatedSprite) constrain() {
	switch as.direction.Y {
	// horizontal check
	case 0:
		if as.rect.X < as.pathRect.Left() ||
			as.rect.X > as.pathRect.Right() {
			as.direction.X *= -1
		}
	default:
		if as.rect.Y < as.pathRect.Top() ||
			as.rect.Y > as.pathRect.Bottom() {
			as.direction.Y *= -1
		}
	}
}
