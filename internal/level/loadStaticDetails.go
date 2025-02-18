package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadStaticDetails(objs []*tiled.Object) error {
	for _, obj := range objs {
		var s *sprites.BasicSprite
		var err error
		if obj.Name == "floor_spikes" {
			s, err = sprites.NewSprite(
				obj.Name, rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets,
				sprites.WithImgPos(rl.NewVector2(0, 32)),
				sprites.WithImgHeight(32),
			)
			if err != nil {
				return err
			}
			if obj.Properties.GetBool("inverted") {
				s.FlipV()
			}
		} else {
			s, err = sprites.NewSprite(
				obj.Name, rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets,
				sprites.WithImgWidth(float32(obj.Width)),
				sprites.WithImgHeight(float32(obj.Height)),
			)
			if err != nil {
				return err
			}
		}
		switch obj.Name {
		case "floor_spikes":
			l.AddSpriteGroup(s, "all", "damage")
		case "crate":
			l.AddSpriteGroup(s, "all", "collision")
		default:
			l.AddSpriteGroup(s, "all")
		}
	}
	return nil
}
