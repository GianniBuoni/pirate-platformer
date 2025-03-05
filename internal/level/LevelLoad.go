package level

import (
	"encoding/json"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

func (l *Level) Load(mapPath string) error {
	lLoaders := loaders.LevelLoaders()
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return err
	}

	// data unmarshaling
	ld := LevelData{}
	err = json.Unmarshal(data, &ld)
	if err != nil {
		return err
	}
	for _, layer := range ld.Layers {
		if len(layer.Data) > 0 {
			l.loadTiles(ld, layer.Data, layer.Name, lLoaders.Tiles)
		}
	}

	return nil
}
