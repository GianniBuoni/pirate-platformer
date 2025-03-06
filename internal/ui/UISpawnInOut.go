package ui

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

func (ui *UI) updateHearts() error {
	renderedHearts, err := ui.groups.GetIDs("heart")
	if err != nil {
		return err
	}
	renderedHearts = renderedHearts[1:]

	if len(renderedHearts) < ui.stats.PlayerHP() {
		for i := len(renderedHearts); i < ui.stats.PlayerHP(); i++ {
			ui.spawnHeart(i)
		}
	}
	if len(renderedHearts) > ui.stats.PlayerHP() {
		lastId := renderedHearts[len(renderedHearts)-1]
		s, ok := ui.groups.Sprites[lastId]
		if !ok {
			return DeletedError("heart", lastId)
		}
		s.GetID().Kill = true
		ui.spawnParticle(s)
	}

	// refetch hearts before update
	renderedHearts, err = ui.groups.GetIDs("heart")
	if err != nil {
		return err
	}
	renderedHearts = renderedHearts[1:]

	for _, id := range renderedHearts {
		s, ok := ui.groups.Sprites[id]
		if !ok {
			return DeletedError("heart", id)
		}
		err := s.Update()
		if err != nil {
			return err
		}
	}
	return nil
}
