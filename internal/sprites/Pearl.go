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
	Groups   SpriteGroup
	Lifetime time.Duration
}

func NewPearl(o Object, aLib AssetLibrary, a *Assets) (*Pearl, error) {
	p := &Pearl{
		Pos:      newPos(o, a),
		Movement: newMovement(o),
		Lifetime: time.Duration(o.Properties.Lifetime) * time.Second,
	}
	var err error
	p.ID, err = newId(o, aLib, a)
	if err != nil {
		return nil, err
	}
	p.rect.Set(
		Center(o.X+o.Width*p.direction.X, o.Y),
	)
	return p, nil
}

func (p *Pearl) Update() error {
	dt := rl.GetFrameTime()
	p.MoveX(p.rect, dt)
	p.collision()
	return nil
}

func (p *Pearl) Draw(src rl.Texture2D, pos *Pos) {
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, p.rect.Width, p.rect.Height),
		rl.Rectangle(*p.rect),
		rl.Vector2{}, 0, rl.White,
	)
}

func (p *Pearl) collision() error {
	sprites, err := p.Groups.GetSpritesName("collision")
	if err != nil {
		return err
	}
	for _, s := range sprites {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.hitbox), rl.Rectangle(*s.HitBox()),
		) {
			p.GetID().Kill = true
		}
	}
	return nil
}
