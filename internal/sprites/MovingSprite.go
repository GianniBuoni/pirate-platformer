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

func NewMovingSprite(
	obj Object, aLib AssetLibrary, a *Assets,
) (Sprite, error) {
	ms := MovingSprite{
		Pos:       newPos(obj, a),
		Movement:  newMovement(obj),
		Animation: newAnimation(),
	}
	var err error
	ms.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return &ms, nil
}

func (ms *MovingSprite) GetPath(src map[int]*Rect) {
	path, ok := src[ms.GID]
	if !ok {
		return
	}
	ms.pathRect = path
}

func (ms *MovingSprite) Update() error {
	dt := rl.GetFrameTime()
	ms.oldRect.Copy(ms.hitbox)
	ms.MoveX(ms.hitbox, dt)
	ms.MoveY(ms.hitbox, dt)
	if ms.pathRect != nil {
		ms.FlipH = ms.Movement.PathConstrain(ms.hitbox, ms.FlipH)
	}
	ms.Pos.Update()
	ms.animate(ms.rect, ms.Src)
	return nil
}
