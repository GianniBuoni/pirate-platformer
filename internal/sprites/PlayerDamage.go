package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) damageCollision() error {
	if !p.state.CheckState(hit) {
		damage, err := p.Groups.GetSpritesName("damage")
		if err != nil {
			return err
		}
		for _, s := range damage {
			if rl.CheckCollisionRecs(
				rl.Rectangle(*p.hitbox), rl.Rectangle(*s.HitBox()),
			) {
				switch sprite := s.(type) {
				case *Pearl:
					if p.state.CheckState(attack) {
						sprite.direction.X *= -1
						break
					}
					p.setHit()
					sprite.GetID().Kill = true
				case *MovingSprite:
					if sprite.Image == "tooth" && p.state.CheckState(attack) {
						sprite.direction.X *= -1
						sprite.FlipH *= -1
						break
					}
					p.setHit()
				default:
					p.setHit()
				}
			}
		}
	}
	return nil
}

func (p *Player) setHit() {
	p.frameIndex = 0
	p.state.ToggleState(hit, true)
	p.stats.AddHP(-1)
}
