package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadMoving(objs []*tiled.Object) error {
	for _, obj := range objs {
		// dermine move axis for object
		dir := rl.Vector2{}
		if obj.Width > obj.Height {
			dir.X = 1
		} else {
			dir.Y = 1
		}

		// init sprite
		width := float32(obj.Properties.GetInt("width"))
		height := float32(obj.Properties.GetInt("height"))
		s, err := sprites.NewAnimatedSprite(
			obj.Name,
			rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(width), sprites.WithImgHeight(height),
			sprites.WithDirection(dir),
			sprites.WithSpeed(float32(obj.Properties.GetInt("speed"))),
		)
		if err != nil {
			return err
		}

		// define movement pattern (radial or ortholinear)
		if obj.Properties.GetInt("radius") > 0 {
			s.SetRadialMove(
				float64(obj.Properties.GetInt("radius")),
				float64(obj.Properties.GetInt("end_angle")),
			)
		} else {
			path := NewRectangle(
				float32(obj.X), float32(obj.Y),
				float32(obj.Width), float32(obj.Height),
			)
			s.SetPath(path)
		}

		// init groups
		groups := []string{"all", "moving"}

		// add extra options/hitboxes
		switch obj.Name {
		case "saw":
			groups = append(groups, "damage")
		case "spike":
			groups = append(groups, "damage")
			s.SetHitbox(rl.NewVector2(10, 10), width-20, height-20)
		case "helicopter":
			groups = append(groups, "platform")
			s.SetHitbox(
				rl.NewVector2(PlatHitbox, 0), width-PlatHitbox*2, height,
			)
		}

		// add groups
		l.AddSpriteGroup(s, groups...)
	}
	return nil
}
