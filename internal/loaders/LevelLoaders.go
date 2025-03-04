package loaders

func LevelLoaders() map[string]Loader {
	loaders := []Loader{
		&objectLoader,
	}
	loaderMap := map[string]Loader{}
	for _, l := range loaders {
		loaderMap[l.Key()] = l
	}
	return loaderMap
}
