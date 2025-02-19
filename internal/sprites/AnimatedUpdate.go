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
	as.rect.Set(lib.Center(
		as.hitbox.Center().X, as.hitbox.Center().Y,
	))
}

func (as *AnimatedSprite) orthoMove() {
	as.Movement(as.hitbox)
	as.orthoConstrain()
}

func (as *AnimatedSprite) radialMove() {
	dt := rl.GetFrameTime()
	// load moving sets direction.Y to 1 by default
	as.angle += float64(as.direction.Y * as.speed * dt)
	as.radialConstrain()

	// calc center point of sprite
	radian := as.angle * math.Pi / 180
	x := float64(as.oldRect.Center().X) + math.Cos(radian)*as.radius
	y := float64(as.oldRect.Center().Y) + math.Sin(radian)*as.radius

	// update hitbox
	as.hitbox.Set(lib.Center(float32(x), float32(y)))
}

func (as *AnimatedSprite) orthoConstrain() {
	switch as.direction.Y {
	// horizontal check
	case 0:
		if as.hitbox.X <= as.pathRect.Left() {
			as.direction.X = 1
			as.flipH = 1
		}
		if as.hitbox.X >= as.pathRect.Right() {
			as.direction.X = -1
			as.flipH = -1
		}
	default:
		if as.hitbox.Y <= as.pathRect.Top() {
			as.direction.Y = 1
		}
		if as.hitbox.Y >= as.pathRect.Bottom() {
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
