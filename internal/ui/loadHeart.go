package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var heartLoader = SpriteLoader[Object]{
	Key:     "heart",
	Builder: ObjectMiddleWare(NewAnimatedSprite),
	Groups:  []string{"heart"},
}

func (ui *UI) spawnHeart(i int) error {
	obj, err := ui.assets.GetObject("heart")
	if err != nil {
		return err
	}
	startPos := rl.NewVector2(32, 24)
	newPos := rl.NewVector2(
		startPos.X+(float32(i)*(TextSpacing+obj.Width)),
		startPos.Y,
	)
	obj.X = newPos.X
	obj.Y = newPos.Y
	err = heartLoader.Run(obj, UiLib, ui)
	if err != nil {
		return err
	}
	return nil
}
