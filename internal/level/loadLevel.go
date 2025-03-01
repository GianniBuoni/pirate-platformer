package level

import (
	"fmt"
)

func (l *LevelData) Load(loader *Loaders) error {
	for _, props := range loader.props {
		err := props.Run(l.MapProps, l)
		if err != nil {
			return err
		}
		continue
	}
	for _, layer := range l.Layers {
		if len(layer.Data) > 0 {
			loader, ok := loader.tiles[layer.Name]
			if !ok {
				fmt.Printf("loader %s not yet implemented.\n", layer.Name)
				continue
			}
			err := loader.Run(layer.Data, l)
			if err != nil {
				return err
			}
			continue
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
	return nil
}
