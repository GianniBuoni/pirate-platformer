package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Heart struct {
	ID
	Pos
	Animation
	canAnimate bool
}

func NewHeartSprite(
	o Object, aLib AssetLibrary, a *Assets,
) (Sprite, error) {
	hs := &Heart{
		Pos:        newPos(o, a),
		Animation:  newAnimation(),
		canAnimate: true,
	}
	var err error
	hs.ID, err = newId(o, aLib, a)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (hs *Heart) Update() error {
	switch hs.canAnimate {
	case true:
		hs.animate(hs.rect, hs.Src)
		if int(hs.frameIndex) >= hs.frameCount {
			hs.frameIndex = 0
			hs.canAnimate = false
		}
	case false:
		catylst, err := RandInt(0, 500)
		if err != nil {
			return err
		}
		if catylst == 0 {
			hs.canAnimate = true
		}
	}
	return nil
}
