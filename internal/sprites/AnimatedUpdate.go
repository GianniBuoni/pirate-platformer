package sprites

import (
	"math"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (as *AnimatedSprite) Update() {
	if as.pathRect != nil {
		as.oldRect.Copy(as.hitbox)
		as.orthoMove()
	} else {
		as.radialMove()
	}
}

func (as *AnimatedSprite) orthoMove() {
	as.Movement(as.rect)
	as.orthoConstrain()
}

func (as *AnimatedSprite) radialMove() {
	dt := rl.GetFrameTime()
	// load moving sets direction.Y to 1 by default
	as.angle += float64(as.direction.Y * as.speed * dt)
	as.radialConstrain()
	x := float64(as.oldRect.Center().X) + math.Cos(as.angle)*as.radius
	y := float64(as.oldRect.Center().Y) + math.Sin(as.angle)*as.radius
	as.rect.Set(lib.Center(
		float32(x), float32(y),
	))
}

func (as *AnimatedSprite) orthoConstrain() {
	switch as.direction.Y {
	// horizontal check
	case 0:
		if as.rect.X <= as.pathRect.Left() {
			as.direction.X = 1
			as.flipH = 1
		}
		if as.rect.X >= as.pathRect.Right() {
			as.direction.X = -1
			as.flipH = -1
		}
	default:
		if as.rect.Y <= as.pathRect.Top() {
			as.direction.Y = 1
		}
		if as.rect.Y >= as.pathRect.Bottom() {
			as.direction.Y = -1
		}
	}
}

func (as *AnimatedSprite) radialConstrain() {
	if as.endAngle != 0 {
		if as.angle <= 0 {
			as.direction.Y = 1
		}
		if as.angle >= as.endAngle {
			as.direction.Y = -1
		}
	}
}
