package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionSide uint

const (
	floor CollisionSide = iota
	left
	right
	air
)

func (p *Player) checkCollisionSide() (CollisionSide, error) {
	// check floor
	groupsToCheck := []string{"collision", "platform"}
	for _, group := range groupsToCheck {
		sprites, err := p.Groups.GetSpritesName(group)
		if err != nil {
			return 0, err
		}
		for _, s := range sprites {
			if rl.CheckCollisionRecs(
				rl.Rectangle(*p.cRects[floor]), rl.Rectangle(*s.HitBox()),
			) {
				return floor, nil
			}
		}
	}
	// check left and right
	sides := map[CollisionSide]rl.Rectangle{
		left:  rl.Rectangle(*p.cRects[left]),
		right: rl.Rectangle(*p.cRects[right]),
	}
	sprites, err := p.Groups.GetSpritesName("wall")
	if err != nil {
		return 0, err
	}
	for k, v := range sides {
		for _, s := range sprites {
			if rl.CheckCollisionRecs(rl.Rectangle(v), rl.Rectangle(*s.HitBox())) {
				return k, nil
			}
		}
	}
	return air, nil
}

func (p *Player) getCRects() {
	p.cRects[floor] = NewRectangle(0, 0, p.hitbox.Width, 4)
	p.cRects[left] = NewRectangle(0, 0, 2, p.hitbox.Height-4)
	p.cRects[right] = NewRectangle(0, 0, 2, p.hitbox.Height-4)
}

func (p *Player) updateCRects() {
	p.cRects[floor].Set(
		Center(p.hitbox.Center().X, p.hitbox.Center().Y), Top(p.hitbox.Bottom()),
	)
	p.cRects[left].Set(
		Bottom(p.hitbox.Bottom()), Right(p.hitbox.Left()),
	)
	p.cRects[right].Set(
		Bottom(p.hitbox.Bottom()), Left(p.hitbox.Right()),
	)
}
