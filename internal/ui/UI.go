package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type UI struct {
	groups SpriteGroup
	texts  map[string]Text
	assets *Assets
	stats  *Stats
	nextId int
}

func NewUI(s *Stats, a *Assets) (*UI, error) {
	return &UI{
		groups: SpriteGroup{
			IDs: map[string][]int{
				"ephemeral": {},
			},
			Sprites: map[int]Sprite{},
		},
		texts:  map[string]Text{},
		assets: a,
		stats:  s,
	}, nil
}

func (ui *UI) AddSpriteGroup(s Sprite, groups ...string) {
	// assign ids
	id := ui.NextId()
	s.GetID().GID = id
	// register sprites
	ui.groups.Sprites[id] = s
	// register groups
	for _, group := range groups {
		ui.groups.IDs[group] = append(ui.groups.IDs[group], id)
	}
}

func (ui *UI) Assets() *Assets {
	return ui.assets
}

func (ui *UI) Texts() map[string]Text {
	return ui.texts
}

func (ui *UI) NextId() int {
	id := ui.nextId
	ui.nextId++
	return id
}
