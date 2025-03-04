package loaders

func UILoaders() map[string]Loader {
	loaders := []Loader{
		// Text
		&bodyLeftLoader, &bodyCenterLoader, &displayCenterLoader,
		// Spites
		&HeartLoader, &coinLoader,
	}
	loaderMap := map[string]Loader{}
	for _, l := range loaders {
		loaderMap[l.Key()] = l
	}
	return loaderMap
}
