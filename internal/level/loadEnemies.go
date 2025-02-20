package level

import (
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	"github.com/lafriks/go-tiled"
)

var toothPath *lib.Rect

func (l *LevelData) loadEnemies(objs []*tiled.Object) error {
	for _, obj := range objs {
		if obj.Name == "shell" {
			rect := lib.NewRectangle(
				float32(obj.X), float32(obj.Y), float32(obj.Width), float32(obj.Height),
			)
			flipH := obj.Properties.GetBool("inverted")
			s, err := sprites.NewShell(
				rect, l.levelAssets, flipH,
			)
			if err != nil {
				return err
			}
			l.AddSpriteGroup(s, "all", "collision", "moving", "shell")
		}
		if strings.Contains(obj.Name, "tooth") {
			rect := lib.NewRectangle(
				float32(obj.X), float32(obj.Y),
				float32(obj.Width), float32(obj.Height),
			)
			if obj.Name == "tooth_path" {
				toothPath = rect
			} else {
				s, err := sprites.NewTooth(
					rect, toothPath, l.levelAssets,
				)
				if err != nil {
					return err
				}
				l.AddSpriteGroup(s, "all", "moving", "damage")
			}
		}
	}
	return nil
}
