package sprites

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) damageCollision() {
	if !p.state.CheckState(hit) {
		for _, s := range p.Groups["damage"] {
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
}

func (p *Player) setHit() {
	p.frameIndex = 0
	p.state.ToggleState(hit, true)
	p.stats.AddHP(-1)
	fmt.Printf("PlayerHP: %v\n", p.stats.PlayerHP())
}
