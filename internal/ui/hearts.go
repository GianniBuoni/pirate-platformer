package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/loaders"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (ui *UI) spawnHeart(i int) error {
	// get heart tempate
	obj, err := ui.assets.GetObject("heart")
	if err != nil {
		return err
	}

	// copy pos of the original heart at index 0
	heartId := ui.groups["heart"][0]
	heartPos := ui.sprites[heartId].GetPos()
	startPos := rl.NewVector2(heartPos.Rect().X, heartPos.Rect().Y)

	newPos := rl.NewVector2(
		startPos.X+(float32(i)*(TextSpacing+obj.Width)),
		startPos.Y,
	)

	obj.X = newPos.X
	obj.Y = newPos.Y
	err = loaders.HeartLoader.Run(obj, UiLib, ui)
	if err != nil {
		return err
	}
	return nil
}
