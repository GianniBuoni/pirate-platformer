package sprites

import (
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pearl struct {
	Pos
	ID
	Movement
	Groups   map[string][]Sprite
	Lifetime time.Duration
}

func NewPearl(o Object, a *Assets) (*Pearl, error) {
	id, err := newId(o, ImageLib, a)
	if err != nil {
		return nil, err
	}
	p := &Pearl{
		Pos:      newPos(o, a),
		ID:       id,
		Movement: newMovement(o),
		Lifetime: time.Duration(o.Properties.Lifetime) * time.Second,
	}
	p.rect.Set(
		Center(o.X+o.Width*p.direction.X, o.Y),
	)
	return p, nil
}

func (p *Pearl) Update() {
	dt := rl.GetFrameTime()
	p.MoveX(p.rect, dt)
	p.collision()
}

func (p *Pearl) Draw(id *ID, pos *Pos) error {
	src, err := p.assets.GetImage(p.assetLib, p.Image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, p.rect.Width, p.rect.Height),
		rl.Rectangle(*p.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}

func (p *Pearl) collision() {
	for _, s := range p.Groups["collision"] {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.hitbox), rl.Rectangle(*s.HitBox()),
		) {
			p.GetID().Kill = true
		}
	}
}
