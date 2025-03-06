package ui

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (ui *UI) Update() error {
	// update ephemeral group
	ephemeral, ok := ui.groups["ephemeral"]
	if ok {
		for _, id := range ephemeral {
			s, ok := ui.sprites[id]
			if !ok {
				return DeletedError("ephemeral", id)
			}
			err := s.Update()
			if err != nil {
				return err
			}
		}
	}
	// update heart group
	err := ui.updateHearts()
	if err != nil {
		return err
	}
	ui.spriteCleanup("heart", "ephemeral")
	return nil
}
