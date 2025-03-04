package level

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

func (l *LevelData) Load(mapPath string) error {
	// init loaders
	lLoaders := loaders.LevelLoaders()

	// new level structs
	layers := Layers{}
	t := TileRefs{}

	data, err := os.ReadFile(mapPath)
	if err != nil {
		return err
	}

	// unmarshal level data
	err = json.Unmarshal(data, l)
	if err != nil {
		return err
	}
	// unmarshal layer data
	err = json.Unmarshal(data, &layers)
	if err != nil {
		return err
	}
	// unmarshal tile ref data
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	// parse tile refs
	for _, ref := range t.TileRef {
		name := GetAssetKey(ref.Source)
		tileset, ok := l.assets.TilesetData[name]
		if !ok {
			return fmt.Errorf(
				"key: %s\n not found in levelAssets tileset data.", name,
			)
		}
		key := GIDRange{
			FirstGID: ref.FirstGID,
			LastGID:  ref.FirstGID + tileset.Count - 1,
		}
		l.tileRefs[key] = name
	}

	// parse layers
	for _, layer := range layers.Layers {
		if len(layer.Data) > 0 {
			fmt.Println("tile layer")
			continue
		}
		for _, obj := range layer.Objects {
			key := obj.Properties.Loader
			lo, ok := lLoaders[key]
			if !ok {
				fmt.Printf("Loader \"%s\" not impemented.\n", key)
				continue
			}
			lo.Run(obj, ImageLib, l)
		}
	}
	return nil
}
