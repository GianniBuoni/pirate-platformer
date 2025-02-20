package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadItems(objs []*tiled.Object) error {
	for _, item := range objs {
		pos := rl.NewVector2(float32(item.X), float32(item.Y))
		s, err := sprites.NewAnimatedSprite(
			item.Name, pos, l.levelAssets,
		)
		if err != nil {
			return err
		}
		l.AddSpriteGroup(s, "all")
	}
	return nil
}
