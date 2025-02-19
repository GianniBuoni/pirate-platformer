package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadObjects(objs []*tiled.Object) error {
	for _, obj := range objs {
		// init sprite
		width := float32(obj.Width)
		height := float32(obj.Height)
		s, err := sprites.NewAnimatedSprite(
			obj.Name,
			rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(width), sprites.WithImgHeight(height),
		)
		if err != nil {
			return err
		}

		// groups
		groups := []string{"all"}

		// other options
		switch obj.Name {
		case "flag":
			s.SetHitbox(rl.NewVector2(36, 0), 20, height)
		case "palm":
			speed, err := RandInt(-2, 2)
			if err != nil {
				return err
			}
			s.SetFrameSpeed(FrameSpeed + float32(speed))
			s.SetHitbox(rl.NewVector2(20, 0), width-40, height)
			groups = append(groups, "platform")
		case "floor_spikes":
			if obj.Properties.GetBool("inverted") {
				s.FlipV()
			}
			groups = append(groups, "damage")
		}
		l.AddSpriteGroup(s, groups...)
	}
	return nil
}
