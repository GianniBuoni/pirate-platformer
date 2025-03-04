package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

var bodyLeftLoader = TextLoader{
	key:     "bodyLeft",
	Builder: textMiddleWare(NewText),
}

var displayCenterLoader = TextLoader{
	key:     "displayCenter",
	Builder: textMiddleWare(NewText),
}

var bodyCenterLoader = TextLoader{
	key:     "bodyCenter",
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
