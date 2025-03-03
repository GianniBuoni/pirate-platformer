package ui

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

func newUiLoaders() map[string]SpriteLoader[Object] {
	loaders := []SpriteLoader[Object]{
		ObjectLoader,
	}
	uiLoaders := map[string]SpriteLoader[Object]{}
	for _, loader := range loaders {
		uiLoaders[loader.Key] = loader
	}
	return uiLoaders
}

func (ui *UI) Load(mapPath string) error {
	loaderMap := newUiLoaders()

	// find get map objects and make sprites
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return err
	}
	layerData := Layers{}
	err = json.Unmarshal(data, &layerData)
	if err != nil {
		return err
	}

	// proess unmarshaled layers
	for _, layer := range layerData.Layers {
		if len(layer.Objects) == 0 {
			continue
		}
		for _, obj := range layer.Objects {
			loadKey := obj.Properties.Loader
			if loadKey == "" {
				fmt.Printf(
					"obj \"%s\" does not have a loader assigned in Tiled\n",
					obj.Image,
				)
				continue
			}
			loader, ok := loaderMap[loadKey]
			if !ok {
				fmt.Printf("loader \"%s\" not implemented\n", loadKey)
				continue
			}
			err := loader.Run(obj, UiLib, ui)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
