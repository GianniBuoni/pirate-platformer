package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *LevelData) Update() error {
	return nil
}

func (l *LevelData) checkShells() error {
	for _, id := range l.groups["shell"] {
		s := l.sprites[id]
		shell, ok := s.(*Shell)
		if !ok {
			return fmt.Errorf(
				"Sprite %s, in sprite group \"shell\" is not a shell sprite",
				s.GetID().Image,
			)
		}
		if shell.SpawnFrame() {
			// TODO spawnPearl
			shell.Attack = false
			return nil
		}
	}
	return nil
}

func (l *LevelData) checkPearls() {
}

func (l *LevelData) checkClouds() {
}

func (l *LevelData) cleanup(groups ...string) {
}
