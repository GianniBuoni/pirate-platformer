package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *Level) spawnInOut() error {
	// check shells
	ids, err := l.groups.GetIDs("shell")
	if err != nil {
		return err
	}
	shells, err := l.groups.GetSprites(ids, "shell")
	for _, s := range shells {
		shell, ok := s.(*sprites.Shell)
		if !ok {
			continue
		}
		if shell.SpawnFrame() {
			l.spawnPearl(shell)
			shell.Attack = false
		}
	}
	// check pearls
	ids, err = l.groups.GetIDs("ephemeral")
	if err != nil {
		return err
	}
	ephemeral, err := l.groups.GetSprites(ids, "ephemeral")
	if err != nil {
		return err
	}
	for _, s := range ephemeral {
		pearl, ok := s.(*sprites.Pearl)
		if !ok {
			continue
		}
		if pearl.Kill {
			l.spawnParticle(pearl)
		}
	}
	return nil
}
