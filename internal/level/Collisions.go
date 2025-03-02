package level

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *LevelData) itemCollisions() {
	for _, sprite := range l.groups["item"] {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*sprite.HitBox()), rl.Rectangle(*l.player.HitBox()),
		) {
			if !sprite.GetID().Kill {
				item := sprite.(*sprites.Item)
				switch item.GetID().Image {
				case "gold", "silver", "diamond":
					l.stats.Coins += item.Value
					fmt.Printf("Coins: %v\n", l.stats.Coins)
				case "skull":
					l.stats.SetMaxHP(item.Value)
					fmt.Println("MaxHP")
					fallthrough
				default:
					l.stats.AddHP(item.Value)
					fmt.Printf("PlayerHP: %v\n", l.stats.PlayerHP())
				}
				item.GetID().Kill = true
				l.spawnParticle(item)
			}
		}
	}
}
