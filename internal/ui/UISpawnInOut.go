package ui

import (
	"errors"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (ui *UI) fetchHearts() ([]int, error) {
	renderedHearts, ok := ui.groups["heart"]
	if !ok {
		return nil, errors.New(
			"Oops! The original heart spirte might have been deleted.",
		)
	}
	renderedHearts = renderedHearts[1:]
	return renderedHearts, nil
}

func (ui *UI) updateHearts() error {
	renderedHearts, err := ui.fetchHearts()
	if err != nil {
		return err
	}
	if len(renderedHearts) < ui.stats.PlayerHP() {
		for i := len(renderedHearts); i < ui.stats.PlayerHP(); i++ {
			ui.spawnHeart(i)
		}
	}
	if len(renderedHearts) > ui.stats.PlayerHP() {
		lastId := renderedHearts[len(renderedHearts)-1]
		s, ok := ui.sprites[lastId]
		if !ok {
			return DeletedError("heart", lastId)
		}
		s.GetID().Kill = true
		ui.spawnParticle(s)
	}
	// refetch hearts before update
	renderedHearts, err = ui.fetchHearts()
	if err != nil {
		return err
	}
	for _, id := range renderedHearts {
		s, ok := ui.sprites[id]
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

func (ui *UI) spriteCleanup(groups ...string) error {
	for _, name := range groups {
		group, ok := ui.groups[name]
		if !ok {
			continue
		}
		for i, id := range group {
			s, ok := ui.sprites[id]
			if !ok {
				return DeletedError(name, id)
			}
			if s.GetID().Kill {
				ui.groups[name] = removeSliceIndex(i, group)
				delete(ui.sprites, id)
			}
		}
	}
	return nil
}

func removeSliceIndex(i int, src []int) []int {
	last := len(src) - 1
	src[i] = src[last]
	return src[:last]
}
