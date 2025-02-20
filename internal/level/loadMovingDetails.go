package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadMovingDetails(objs []*tiled.Object) error {
	for _, obj := range objs {
		// dermine move axis for object
		dir := rl.Vector2{}
		if obj.Width > obj.Height {
			dir.X = 1
		} else {
			dir.Y = 1
		}

		switch obj.Name {
		case "spike":
			// spke chain info
			img := "spike_chain"
			src, err := l.levelAssets.GetImage(ImageLib, img)
			if err != nil {
				return err
			}
			width, height := float32(src.Height), float32(src.Width)
			centerX, centerY := float32(obj.X)+(lib.TileSize/2), float32(obj.Y)+(lib.TileSize/2)
			imgX, imgY := centerX-(width/2), centerY-(height/2)

			maxRadius := float64(obj.Properties.GetInt("radius"))
			endAngle := float64(obj.Properties.GetInt("end_angle"))

			for r := float64(0); r < maxRadius; r += 20 {
				s, err := sprites.NewAnimatedSprite(
					img, rl.NewVector2(imgX, imgY), l.levelAssets,
					sprites.WithImgWidth(width),
					sprites.WithImgHeight(height),
					sprites.WithDirection(dir),
					sprites.WithSpeed(float32(obj.Properties.GetInt("speed"))),
				)
				if err != nil {
					return err
				}
				s.SetRadialMove(r, endAngle)
				l.AddSpriteGroup(s, "all", "moving")
			}
		case "saw":
			img := "saw_chain"
			src, err := l.levelAssets.GetImage(ImageLib, img)
			if err != nil {
				return err
			}

			width, height := float32(src.Width), float32(src.Height)
			sawOffset := float32(obj.Properties.GetInt("width") / 2)
			var (
				max float32
				pos rl.Vector2
			)

			if dir.X == 1 {
				max = float32(obj.Width)
			} else {
				max = float32(obj.Height)
			}

			for i := float32(0); i < max; i += 20 {
				if dir.X == 1 {
					pos = rl.NewVector2(float32(obj.X)+i, float32(obj.Y)+sawOffset)
				} else {
					pos = rl.NewVector2(float32(obj.X)+sawOffset, float32(obj.Y)+i)
				}
				s, err := sprites.NewSprite(
					"saw_chain", pos, l.levelAssets,
					sprites.WithImgWidth(width), sprites.WithImgHeight(height),
				)
				if err != nil {
					return err
				}
				l.AddSpriteGroup(s, "all")
			}
		}
	}
	return nil
}
