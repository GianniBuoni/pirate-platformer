package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
)

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
