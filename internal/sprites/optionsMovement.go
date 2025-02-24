package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Movement struct {
	pathRect  *Rect
	direction rl.Vector2
	speed     float32
	gravity   float32
}

func newMovement(obj Object) Movement {
	return Movement{
		direction: rl.Vector2{
			X: obj.Properties.DirX,
			Y: obj.Properties.DirY,
		},
		speed: obj.Properties.Speed,
	}
}

func (m *Movement) MoveX(r *Rect, dt float32) {
	r.X += m.direction.X * m.speed * dt
}

func (m *Movement) MoveY(r *Rect, dt float32) bool {
	m.direction.Y += m.gravity * dt
	r.Y += m.direction.Y * m.speed * dt
	if m.direction.Y != 0 {
		return true
	}
	return false
}

func (m *Movement) SetGravity(b bool, multiplier float32) {
	switch b {
	case true:
		m.gravity = Gravity
	default:
		m.direction.Y = 0
		m.gravity = Gravity * multiplier
	}
}

func (m *Movement) PathConstrain(r *Rect, flipH float32) float32 {
	switch m.direction.Y {
	// horizontal check
	case 0:
		if r.Left() <= m.pathRect.Left() {
			m.direction.X = 1
			return 1
		}
		if r.Right() >= m.pathRect.Right() {
			m.direction.X = -1
			return -1
		}
		// vertical check
	default:
		if r.Top() <= m.pathRect.Top() {
			m.direction.Y = 1
		}
		if r.Bottom() >= m.pathRect.Bottom() {
			m.direction.Y = -1
		}
	}
	return flipH
}
