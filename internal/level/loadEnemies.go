package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	"github.com/lafriks/go-tiled"
)

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
	}
	return nil
}
