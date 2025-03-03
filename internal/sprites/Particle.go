package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Particle struct {
	Pos
	ID
	Animation
}

func NewParticle(o Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	p := &Particle{
		Pos:       newPos(o, a),
		Animation: newAnimation(),
	}
	var err error
	p.ID, err = newId(o, aLib, a)
	if err != nil {
		return nil, err
	}
	p.rect.Set(Center(o.X, o.Y))
	p.animate(p.rect, p.Src)
	return p, nil
}

func (p *Particle) Update() error {
	p.animate(p.rect, p.Src)
	if int(p.frameIndex) >= p.frameCount {
		p.GetID().Kill = true
	}
	return nil
}
