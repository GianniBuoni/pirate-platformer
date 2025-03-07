package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) collison(axis string) error {
	sprites, err := p.Groups.GetSpritesName("collision")
	if err != nil {
		return err
	}
	for _, s := range sprites {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*s.HitBox()), rl.Rectangle(*p.hitbox),
		) {
			switch axis {
			case "x":
				if p.hitbox.Left() <= s.HitBox().Right() &&
					p.oldRect.Left() >= s.OldRect().Right() {
					p.hitbox.Set(Left(s.HitBox().Right()))
				}
				if p.hitbox.Right() >= s.HitBox().Left() &&
					p.oldRect.Right() <= s.OldRect().Left() {
					p.hitbox.Set(Right(s.HitBox().Left()))
				}
			case "y":
				if p.hitbox.Top() <= s.HitBox().Bottom() &&
					p.oldRect.Top() >= s.OldRect().Bottom() {
					p.hitbox.Set(Top(s.HitBox().Bottom()))
				}
				if p.hitbox.Bottom() >= s.HitBox().Top() &&
					p.oldRect.Bottom() <= s.OldRect().Top() {
					p.hitbox.Set(Bottom(s.HitBox().Top()))
				}
				p.direction.Y = 0
			}
		}
	}
	return nil
}

func (p *Player) patformCollision() error {
	sprites, err := p.Groups.GetSpritesName("platform")
	if err != nil {
		return err
	}
	for _, pl := range sprites {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*p.hitbox), rl.Rectangle(*pl.HitBox()),
		) && p.state.CheckState(canPlatform) {
			if p.hitbox.Bottom() >= pl.HitBox().Top() &&
				p.oldRect.Bottom() <= pl.OldRect().Top() {
				p.hitbox.Set(Bottom(pl.HitBox().Top()))
				p.direction.Y = 0
				p.platform = pl
			}
		}
	}
	return nil
}
