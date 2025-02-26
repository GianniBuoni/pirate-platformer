package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type AnimatedSprite struct {
	ID
	Pos
	Animation
}

func NewAnimatedSprite(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	as := AnimatedSprite{
		ID:        id,
		Pos:       newPos(obj, a),
		Animation: newAnimation(),
	}
	return &as, nil
}

func (as *AnimatedSprite) Update() {}
