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
				fmt.Printf("%s: %v\n", item.Name(), item.Value)
				item.GetID().Kill = true
				l.spawnParticle(item)
			}
		}
	}
}
