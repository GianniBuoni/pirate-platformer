package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

func (l *Level) Load(ld LevelData) error {
	// init loaders
	lLoaders := loaders.LevelLoaders()
	l.addLocalLoaders(lLoaders.Objects)

	ld.MapProps.Width = ld.Width
	ld.MapProps.Height = ld.Height
	l.nextLevel = ld.MapProps.NextLevel
	l.Top = float32(ld.MapProps.TopLimit) * TileSize
	l.Width = float32(ld.Width) * TileSize
	l.Height = float32(ld.Height) * TileSize

	// load sprites
	for _, mpLoader := range lLoaders.MapProps {
		err := l.loadMapProps(ld.MapProps, mpLoader)
		if err != nil {
			return err
		}
	}
	for _, layer := range ld.Layers {
		if len(layer.Data) > 0 {
			err := l.loadTiles(ld, layer.Data, layer.Name, lLoaders.Tiles)
			if err != nil {
				return err
			}
			continue
		}
		for _, o := range layer.Objects {
			err := l.loadObjects(o, lLoaders.Objects)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *Level) addLocalLoaders(ldOject map[string]loaders.Loader[Object]) {
	local := []loaders.Loader[Object]{
		&platformLoader, &pathLoader, &sawPathLoader, &damageLoader, &playerLoader,
		&shellLoader,
	}
	for _, loader := range local {
		ldOject[loader.GetKey()] = loader
	}
}
