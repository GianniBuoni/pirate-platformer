package level

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *Level) itemCollisions() error {
	items, err := l.groups.GetSpritesName("item")
	if err != nil {
		return err
	}
	for _, s := range items {
		if rl.CheckCollisionRecs(
			rl.Rectangle(*s.HitBox()), rl.Rectangle(*l.player.HitBox()),
		) {
			if !s.GetID().Kill {
				item, ok := s.(*sprites.Item)
				if !ok {
					return fmt.Errorf(
						"type mismatch, \"%s: %d\" is in the item sprite group",
						s.GetID().Image, s.GetID().GID,
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

func (l *Level) checkDeathVictory() error {
	if l.player.HitBox().Top() >= l.Height {
		l.stats.AddHP(-l.stats.PlayerHP())
	}
	flags, err := l.groups.GetSpritesName("flag")
	if err != nil {
		return err
	}
	if rl.CheckCollisionRecs(
		rl.Rectangle(*l.player.HitBox()), rl.Rectangle(*flags[0].HitBox()),
	) {
		l.stats.SetVictory(l.nextLevel)
	}
	return nil
}
