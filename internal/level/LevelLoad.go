package level

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) Load() error {
	layerFuncs := map[string]func([]*tiled.Object) error{
		"static details": l.loadStaticDetails,
		"objects":        l.loadObjects,
		"items":          l.loadItems,
		"enemies":        l.loadEnemies,
	}
	for _, objGroup := range l.mapData.ObjectGroups {
		switch objGroup.Name {
		case "zero index":
			if err := l.loadZero(objGroup.Objects); err != nil {
				return err
			}
			if err := l.loadTiles(); err != nil {
				return err
			}
		case "moving":
			if err := l.loadMovingDetails(objGroup.Objects); err != nil {
				return err
			}
			if err := l.loadMoving(objGroup.Objects); err != nil {
				return err
			}
		case "player":
			if objGroup.Name == "player" {
				if err := l.loadPlayer(objGroup.Objects[0]); err != nil {
					return err
				}
			}
		default:
			loadFunc, ok := layerFuncs[objGroup.Name]
			if ok {
				if err := loadFunc(objGroup.Objects); err != nil {
					return err
				}
			} else {
				fmt.Printf("objGroup.Name: %s not yet implemented\n", objGroup.Name)
			}
		}
	}

	for _, sprite := range l.groups["shell"] {
		shell, ok := sprite.(*sprites.ShellSprite)
		if ok {
			shell.SetPlayer(l.player)
		}
	}
	return nil
}
