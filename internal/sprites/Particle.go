package sprites

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	Pos
	ID
	Movement
	Animation
	Lifetime time.Duration
}

func NewParticle(image string, s Sprite, a *Assets) (*Particle, error) {
	src, err := a.GetImage(ImageLib, image)
	if err != nil {
		return nil, err
	}
	id := ID{
		image:    image,
		assets:   a,
		assetLib: ImageLib,
	}
	obj := Object{
		Properties: Properties{
			DirX:   s.GetPos().flipH,
			SpeedX: 128,
		},
		Width:  float32(src.Width),
		Height: float32(src.Height),
	}
	p := Particle{
		Pos:      newPos(obj, a),
		ID:       id,
		Movement: newMovement(obj),
	}
	p.rect.Set(Center(
		s.HitBox().Center().X, s.HitBox().Center().Y,
	))
	return &p, nil
}

func NewPearl(s Sprite, a *Assets) (*Particle, error) {
	p, err := NewParticle("pearl", s, a)
	if err != nil {
		return nil, err
	}
	if p.direction.X > 0 {
		p.rect.Set(Center(
			s.HitBox().Right(), s.HitBox().Center().Y,
		))
	} else {
		p.rect.Set(Center(
			s.HitBox().Left(), s.HitBox().Center().Y,
		))
	}
	p.Lifetime = 10 * time.Second
	return p, nil
}

func (p *Particle) Update() {
	dt := rl.GetFrameTime()
	p.MoveX(p.rect, dt)
}
