package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadObjects(objs []*tiled.Object) error {
	for _, obj := range objs {
		s, err := sprites.NewAnimatedSprite(
			obj.Name,
			rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(float32(obj.Width)),
			sprites.WithImgHeight(float32(obj.Height)),
		)
		switch obj.Name {
		case "palm":
			speed, err := RandInt(-2, 2)
			if err != nil {
				return err
			}
			s.SetFrameSpeed(FrameSpeed + float32(speed))
			l.AddSpriteGroup(s, "platform")
		}
		if err != nil {
			return err
		}
		l.AddSpriteGroup(s, "all")
	}
	return nil
}
