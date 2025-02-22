package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Loader[T any] struct {
	Run func(T, *LevelData) error
}

var tileLoaders = map[string]Loader[[]int]{
	"bg":        bgTileLoader,
	"collision": cTileLoader,
	"platform":  pTileLoader,
}

var objectLoaders = map[string]Loader[Object]{
	"object": objectLoader,
}

func (l *LevelData) Load() error {
	for _, layer := range l.Layers {
		if len(layer.Data) > 0 {
			fmt.Printf("loading tiles: %s\n", layer.Name)
			loader := tileLoaders["bg"]
			err := loader.Run(layer.Data, l)
			if err != nil {
				return err
			}
		}
		for _, obj := range layer.Objects {
			loadKey := obj.Properties.Loader
			if loadKey == "" {
				fmt.Printf(
					"obj '%s' does not yet have a loader assigned in tiled\n",
					obj.Image,
				)
				continue
			}
			loader, ok := objectLoaders[loadKey]
			if !ok {
				fmt.Printf("loader: '%s' is not yet implemented\n", loadKey)
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
