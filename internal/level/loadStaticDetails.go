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
		s, err = sprites.NewSprite(
			obj.Name, rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(float32(obj.Width)),
			sprites.WithImgHeight(float32(obj.Height)),
		)
		if err != nil {
			return err
		}
		groups := []string{"all"}
		if obj.Name == "crate" {
			groups = append(groups, "collision")
		}
		l.AddSpriteGroup(s, groups...)
	}
	return nil
}
