package level

import (
	"fmt"

	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *Level) spawnInOut() error {
	for _, id := range l.groups["shell"] {
		s, ok := l.spirtes[id]
		if !ok {
			continue
		}
		shell, ok := s.(*sprites.Shell)
		if !ok {
			return fmt.Errorf(
				"sprite \"%s\" is in the shell sprite group.",
				s.GetID().Image,
			)
		}
		if shell.SpawnFrame() {
			err := l.spawnPearl(shell)
			if err != nil {
				return err
			}
			shell.Attack = false
		}
	}
	l.cleanup("damage", "item")
	return nil
}

func (l *Level) cleanup(groups ...string) {
	for i, id := range l.groups["ephemeral"] {
		s, ok := l.spirtes[id]
		if !ok {
			fmt.Printf("Sprite %d might already be deleted\n", id)
			continue
		}
		if s.GetID().Kill {
			l.groups["ephemeral"] = removeSliceIndex(i, l.groups["ephemeral"])
			delete(l.spirtes, id)
		}
	}
	for _, name := range groups {
		group, ok := l.groups[name]
		if !ok {
			continue
		}
		for i, id := range group {
			_, ok := l.spirtes[id]
			if !ok {
				l.groups[name] = removeSliceIndex(i, group)
			}
		}
	}
}

func removeSliceIndex(i int, src []int) []int {
	last := len(src) - 1
	src[i] = src[last]
	return src[:last]
}
