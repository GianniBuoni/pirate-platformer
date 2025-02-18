package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadMoving(objs []*tiled.Object) error {
	for _, obj := range objs {
		dir := rl.Vector2{}
		if obj.Width > obj.Height {
			dir.X = 1
		} else {
			dir.Y = 1
		}
		s, err := sprites.NewAnimatedSprite(
			obj.Name,
			rl.NewVector2(float32(obj.X), float32(obj.Y)),
			l.levelAssets,
			sprites.WithImgWidth(float32(obj.Properties.GetInt("width"))),
			sprites.WithImgHeight(float32(obj.Properties.GetInt("height"))),
			sprites.WithDirection(dir),
			sprites.WithSpeed(float32(obj.Properties.GetInt("speed"))),
		)
		if err != nil {
			return err
		}
		groups := []string{"all"}
		if obj.Properties.GetBool("platform") {
			groups = append(groups, "platform")
		}
		l.AddSpriteGroup(s, groups...)
	}
	return nil
}
