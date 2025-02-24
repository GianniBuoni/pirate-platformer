package level

import (
	"errors"
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *LevelData) Load(loader *Loaders) error {
	for _, layer := range l.Layers {
		if len(layer.Data) > 0 {
			loader, ok := loader.tiles[layer.Name]
			if !ok {
				fmt.Printf("loader %s not yet implemented.\n", layer.Name)
				continue
			}
			for idx, id := range layer.Data {
				if id == 0 {
					continue
				}
				tile, err := parseTile(idx, id, l)
				if err != nil {
					return err
				}
				err = loader.Run(tile, l)
				if err != nil {
					return err
				}
			}
		}
		for _, obj := range layer.Objects {
			loadKey := obj.Properties.Loader
			if loadKey == "" {
				fmt.Printf(
					"obj '%s' does not have a loader assigned in tiled\n",
					obj.Image,
				)
				continue
			}
			loader, ok := loader.object[loadKey]
			if !ok {
				fmt.Printf("loader '%s' not yet implemented\n", loadKey)
				continue
			}
			err := loader.Run(obj, l)
			if err != nil {
				return err
			}
		}
	}
	// add additional data to player
	playerSprite, ok := l.groups["player"]
	if !ok {
		return errors.New("player not defined in tiled")
	}
	player, ok := playerSprite[0].(*sprites.Player)
	player.Groups = l.groups
	l.player = player
	return nil
}
