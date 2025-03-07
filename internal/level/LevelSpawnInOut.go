package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *Level) spawnInOut() error {
	// check shells
	shells, err := l.groups.GetSpritesName("shell")
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
	ephemeral, err := l.groups.GetSpritesName("ephemeral")
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
