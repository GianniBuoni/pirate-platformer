package level

import (
	"errors"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadObjects(objs []*tiled.Object) error {
	for _, obj := range objs {
		if obj.Name == "palm" {
			s, err := sprites.NewAnimatedSprite(
				"palm", rl.NewVector2(float32(obj.X), float32(obj.Y)),
				l.levelAssets, sprites.WithFrameSize(float32(obj.Width)),
			)
			if err != nil {
				return err
			}
			l.AddSpriteGroup(s, "all", "platform")
		}
		collisonSprites, ok := l.groups["collision"]
		if !ok {
			return errors.New(
				"error Level.loadPlayer(): collision group is not defined.",
			)
		}
		platformSprites, ok := l.groups["platform"]
		if !ok {
			return errors.New(
				"error Level.loadPlayer(): platform group is not defined.",
			)
		}
		if obj.Name == "player" {
			newPlayer := NewPlayerParams{
				Pos:      rl.NewVector2(float32(obj.X), float32(obj.Y)),
				Assets:   l.levelAssets,
				CSprites: &collisonSprites,
				PSprites: &platformSprites,
			}
			sprite, err := sprites.NewPlayer(newPlayer)
			if err != nil {
				return err
			}
			l.AddPlayer(sprite)
		}
	}
	return nil
}
