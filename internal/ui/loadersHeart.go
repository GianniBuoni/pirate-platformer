package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/loaders"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
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
	err = HeartLoader.Run(obj, UiLib, ui)
	if err != nil {
		return err
	}
	return nil
}

func (ui *UI) spawnParticle(s Sprite) error {
	o, err := ui.assets.GetObject("particle")
	if err != nil {
		return err
	}
	o.X = s.HitBox().Center().X
	o.Y = s.HitBox().Center().Y
	err = ParticleLoader.Run(o, ImageLib, ui)
	if err != nil {
		return err
	}
	return nil
}
