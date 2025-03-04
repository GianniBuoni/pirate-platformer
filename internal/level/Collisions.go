package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *LevelData) itemCollisions() {
	for _, id := range l.groups["item"] {
		sprite := l.sprites[id]
		if rl.CheckCollisionRecs(
			rl.Rectangle(*sprite.HitBox()), rl.Rectangle(*l.player.HitBox()),
		) {
			if !sprite.GetID().Kill {
				item := sprite.(*sprites.Item)
				switch item.GetID().Image {
				case "gold", "silver", "diamond":
					l.stats.Coins += item.Value
				case "skull":
					l.stats.SetMaxHP(item.Value)
					fallthrough
				default:
					l.stats.AddHP(item.Value)
				}
				item.GetID().Kill = true
				// TODO spawn particle
			}
		}
	}
}
