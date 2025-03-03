package ui

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	sprites map[string][]Sprite
	texts   map[string]Text
	assets  *Assets
	stats   *Stats
}

func NewUI(s *Stats, a *Assets) (*UI, error) {
	return &UI{
		sprites: map[string][]Sprite{},
		texts:   map[string]Text{},
		assets:  a,
		stats:   s,
	}, nil
}

func (ui *UI) Update() error {
	return nil
}

func (ui *UI) Draw() {
	rl.DrawRectangleRec(
		rl.NewRectangle(0, 0, WindowW, WindowH),
		rl.ColorAlpha(rl.Black, 0.3),
	)
	for _, s := range ui.sprites["all"] {
		s.Draw(s.GetID().Src, s.GetPos())
	}
	coins := fmt.Sprint(ui.stats.Coins)
	ui.texts["coinText"].Draw(coins)
}

func (ui *UI) AddSpriteGroup(s Sprite, groups ...string) {
	for _, group := range groups {
		ui.sprites[group] = append(ui.sprites[group], s)
	}
}

func (ui *UI) Assets() *Assets {
	return ui.assets
}
