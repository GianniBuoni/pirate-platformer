package level

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *Level) itemCollisions() error {
	for _, id := range l.groups["item"] {
		s, ok := l.spirtes[id]
		if !ok {
			fmt.Printf("Sprite \"%d\" might already be deleted", id)
		}
		if rl.CheckCollisionRecs(
			rl.Rectangle(*s.HitBox()), rl.Rectangle(*l.player.HitBox()),
		) {
			if !s.GetID().Kill {
				item, ok := s.(*sprites.Item)
				if !ok {
					return fmt.Errorf(
						"type mismatch, \"%s\" is in the item sprite group",
						s.GetID().Image,
					)
				}
				switch item.Image {
				case "gold", "silver", "diamond":
					l.stats.Coins += item.Value
				case "skull":
					l.stats.SetMaxHP(item.Value)
					fallthrough
				default:
					l.stats.AddHP(item.Value)
				}
				item.GetID().Kill = true
				l.spawnParticle(item)
			}
		}
	}
	return nil
}
