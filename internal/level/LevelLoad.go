package level

import (
	"errors"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *LevelData) Load() error {
	err := l.loadTiles()
	if err != nil {
		return err
	}
	err = l.loadObjects()
	if err != nil {
		return err
	}
	return nil
}

func (l *LevelData) loadObjects() error {
	collisonSprites, ok := l.groups["collision"]
	if !ok {
		return errors.New(
			"error Level.loadPlayer(): collison group is not defined.",
		)
	}
	platformSprites, ok := l.groups["platform"]
	if !ok {
		return errors.New(
			"error Level.loadPlayer(): platform group is not defined.",
		)
	}
	for _, objGroup := range l.mapData.ObjectGroups {
		if objGroup.Name == "BG details" {
			for _, obj := range objGroup.Objects {
				if obj.Type == "static" {
					s, err := sprites.NewSprite(
						obj.Name,
						rl.NewVector2(float32(obj.X), float32(obj.Y)),
						l.levelAssets,
					)
					if err != nil {
						return err
					}
					l.AddSpriteGroup(s, "all")
				}
			}
		}
		if objGroup.Name == "Objects" {
			for _, obj := range objGroup.Objects {
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
		}
	}
	return nil
}
