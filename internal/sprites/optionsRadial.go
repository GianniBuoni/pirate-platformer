package sprites

import (
	"math"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RadialMove struct {
	angle     float64
	endAngle  float64
	radius    float64
	direction float32
	speed     float32
}

func newRadial(obj Object) RadialMove {
	return RadialMove{
		radius:    float64(obj.Properties.DirX),
		direction: obj.Properties.DirY,
		speed:     obj.Properties.SpeedX,
		endAngle:  float64(obj.Properties.SpeedY),
	}
}

func (rm *RadialMove) moveR(src, dest *Rect) {
	rm.angle += float64(rm.direction * rm.speed * rl.GetFrameTime())
	rm.constrainR()

	center := src.Center()
	radian := rm.angle * (math.Pi / 180)
	x := center.X + float32(math.Cos(radian)*rm.radius)
	y := center.Y + float32(math.Sin(radian)*rm.radius)
	dest.Set(Center(x, y))
}

func (rm *RadialMove) constrainR() {
	if rm.endAngle != 0 {
		if rm.angle >= rm.endAngle {
			rm.direction = -1
		}
		if rm.angle <= 0 {
			rm.direction = 1
		}
	}
}
