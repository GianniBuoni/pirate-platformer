package sprites

type Particle struct {
	Pos
	ID
	Animation
}

func (p *Particle) Update() {
	if int(p.frameIndex) >= 3 {
		p.Kill()
	}
}
