package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type UI struct {
	groups  map[string][]int
	sprites map[int]Sprite
	texts   map[string]Text
	assets  *Assets
	stats   *Stats
	nextId  int
}

func NewUI(s *Stats, a *Assets) (*UI, error) {
	return &UI{
		groups:  map[string][]int{},
		sprites: map[int]Sprite{},
		texts:   map[string]Text{},
		assets:  a,
		stats:   s,
	}, nil
}

func (ui *UI) Update() error {
	if len(ui.groups["heart"]) < ui.stats.PlayerHP() {
		for i := range ui.stats.PlayerHP() - len(ui.groups["heart"]) {
			ui.spawnHeart(i)
		}
	}
	for _, int := range ui.groups["heart"] {
		ui.sprites[int].Update()
	}
	return nil
}

func (ui *UI) AddSpriteGroup(
	s Sprite, spriteMap map[int]Sprite, groups ...string,
) {
	spriteMap[s.GetID().GID] = s
	for _, group := range groups {
		ui.groups[group] = append(ui.groups[group], s.GetID().GID)
	}
}

func (ui *UI) Assets() *Assets {
	return ui.assets
}

func (ui *UI) Texts() map[string]Text {
	return ui.texts
}

func (ui *UI) Sprites() map[int]Sprite {
	return ui.sprites
}

func (ui *UI) NextId() int {
	id := ui.nextId
	ui.nextId++
	return id
}
