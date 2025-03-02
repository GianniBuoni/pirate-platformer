package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type UI struct {
	assets *Assets
	stats  *Stats
}

func NewUI(s *Stats, a *Assets) *UI {
	return &UI{
		assets: a,
		stats:  s,
	}
}

func (ui *UI) Update() {}

func (ui *UI) Draw() {}
