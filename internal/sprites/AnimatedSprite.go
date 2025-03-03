package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type AnimatedSprite struct {
	ID
	Pos
	Animation
}

func NewAnimatedSprite(
	obj Object, aLib AssetLibrary, a *Assets,
) (Sprite, error) {
	as := AnimatedSprite{
		Pos:       newPos(obj, a),
		Animation: newAnimation(),
	}
	var err error
	as.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return &as, nil
}

func (as *AnimatedSprite) Update() (err error) {
	as.animate(as.rect, as.Src)
	return nil
}
