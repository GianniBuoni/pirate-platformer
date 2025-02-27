package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Particle struct {
	Pos
	ID
	Animation
}

func NewParticle(o Object, a *Assets) (Sprite, error) {
	id, err := newId(o, ImageLib, a)
	if err != nil {
		return nil, err
	}
	p := &Particle{
		Pos:       newPos(o, a),
		ID:        id,
		Animation: newAnimation(),
	}
	p.rect.Set(Center(o.X, o.Y))
	return p, nil
}

func (p *Particle) Update() {
	if int(p.frameIndex) >= 3 {
		p.Kill()
	}
}
