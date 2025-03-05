package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

func (l *Level) loadMapProps(mp MapProps, loader Loader[MapProps]) error {
	err := loader.Run(mp, TilesetLib, l)
	if err != nil {
		return err
	}
	return nil
}

func (l *Level) loadTiles(
	ld LevelData,
	data []int, key string,
	loaders map[string]Loader[TileParams],
) error {
	gidRanges, err := ld.MapGIDRanges(l.assets)
	if err != nil {
		return err
	}
	tp := TileParams{
		Data:      data,
		GIDRanges: gidRanges,
		Columns:   ld.Width,
	}
	loader, ok := loaders[key]
	if !ok {
		fmt.Printf("Loader \"%s\" not implemented.\n", key)
		return nil
	}
	loader.Run(tp, TilesetLib, l)
	return nil
}

func (l *Level) loadObjects(
	o Object, loaders map[string]Loader[Object],
) error {
	key := o.Properties.Loader
	loader, ok := loaders[key]
	if !ok {
		fmt.Printf("Loader \"%s\" not implemented.\n", key)
		return nil
	}
	err := loader.Run(o, ImageLib, l)
	if err != nil {
		return err
	}
	return nil
}
