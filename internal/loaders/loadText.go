package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

var bodyLeftLoader = TextLoader{
	Key:     "bodyLeft",
	Builder: textMiddleWare(NewText),
}

var displayCenterLoader = TextLoader{
	Key:     "displayCenter",
	Builder: textMiddleWare(NewText),
}

var bodyCenterLoader = TextLoader{
	Key:     "bodyCenter",
	Builder: textMiddleWare(NewText),
}

func textMiddleWare(
	f func(Object) Text,
) func(Object, map[string]Text) {
	return func(o Object, tm map[string]Text) {
		text := f(o)
		tm[o.Image] = text
	}
}
