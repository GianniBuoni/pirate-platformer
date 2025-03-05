package loaders

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

func UILoaders() map[string]Loader[Object] {
	loaders := []Loader[Object]{
		// Text
		&bodyLeftLoader, &bodyCenterLoader, &displayCenterLoader,
		// Spites
		&HeartLoader, &coinLoader,
	}
	loaderMap := map[string]Loader[Object]{}
	for _, l := range loaders {
		loaderMap[l.Key()] = l
	}
	return loaderMap
}
