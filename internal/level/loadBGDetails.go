package level

import (
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadBGDetails(objs []*tiled.Object) error {
	for _, obj := range objs {
		if obj.Name == "candle" {
			s, err := sprites.NewAnimatedSprite(
				"candle_light", rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets,
			)
			if err != nil {
				return err
			}
			l.AddSpriteGroup(s, "all")
		}
		if obj.Type == "static" && !strings.Contains(obj.Name, "palm") {
			// load static sprites
		} else if !strings.Contains(obj.Name, "palm") {
			s, err := sprites.NewAnimatedSprite(
				obj.Name, rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets,
			)
			if err != nil {
				return err
			}
			l.AddSpriteGroup(s, "all")
		}
	}
	return nil
}
