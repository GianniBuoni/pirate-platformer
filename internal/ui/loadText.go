package ui

import (
	"errors"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var loadTextLeft = SpriteLoader[Object]{
	Key:     "textLeft",
	Builder: textMiddleWare(mapText),
}

func textMiddleWare(
	f func(Object, *UI),
) func(Object, AssetLibrary, GameModule) ([]Sprite, error) {
	return func(o Object, al AssetLibrary, gm GameModule) ([]Sprite, error) {
		ui, ok := gm.(*UI)
		if !ok {
			return nil, errors.New("passed in game module is not a the ui.")
		}
		f(o, ui)
		return nil, nil
	}
}

func mapText(o Object, ui *UI) {
	text := Text{
		pos: rl.NewVector2(o.X, o.Y),
	}
	switch o.Properties.Loader {
	case "textCenter":
		text.alignment = center
	}
	ui.texts[o.Image] = text
}
