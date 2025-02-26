package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MSpriteAxis uint

const (
	X MSpriteAxis = iota
	Y
)

type MovingSprite struct {
	Pos
	ID
	Movement
	Animation
}

func NewMovingSprite(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	ms := MovingSprite{
		Pos:       newPos(obj, a),
		ID:        id,
		Movement:  newMovement(obj),
		Animation: newAnimation(),
	}
	return &ms, nil
}

func (ms *MovingSprite) GetPath(src map[int]*Rect) {
	path, ok := src[ms.id]
	if !ok {
		return
	}
	ms.pathRect = path
}

func (ms *MovingSprite) Update() {
	dt := rl.GetFrameTime()
	ms.oldRect.Copy(ms.hitbox)
	ms.MoveX(ms.hitbox, dt)
	ms.MoveY(ms.hitbox, dt)
	if ms.pathRect != nil {
		ms.flipH = ms.Movement.PathConstrain(ms.hitbox, ms.flipH)
	}
	ms.Pos.Update()
}
