package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadStaticDetails(objs []*tiled.Object) error {
	for _, obj := range objs {
		var s *sprites.BasicSprite
		var err error
		width := float32(obj.Width)
		height := float32(obj.Height)
		s, err = sprites.NewSprite(
			obj.Name, rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(width), sprites.WithImgHeight(height),
		)
		if err != nil {
			return err
		}
		groups := []string{"all"}
		switch obj.Name {
		case "crate":
			groups = append(groups, "collision")
			s.SetHitbox(
				rl.NewVector2(lib.PlatHitbox, 0),
				width-lib.PlatHitbox*2, height,
			)
		}
		l.AddSpriteGroup(s, groups...)
	}
	return nil
}
