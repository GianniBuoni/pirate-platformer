package level

import (
	"fmt"
	"time"

	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *LevelData) checkShells() error {
	for _, s := range l.groups["shell"] {
		shell, ok := s.(*Shell)
		if !ok {
			return fmt.Errorf(
				"Sprite %s, in sprite group \"shell\" is not a shell sprite",
				s.Name(),
			)
		}
		if shell.SpawnFrame() {
			err := l.spawnPearl(shell)
			if err != nil {
				return err
			}
			shell.Attack = false
			return nil
		}
	}
	return nil
}

func (l *LevelData) spawnPearl(s Sprite) error {
	p, err := NewPearl(s, l.levelAssets)
	if err != nil {
		return err
	}
	l.AddSpriteGroup(p, "all", "moving", "damage")
	go func() {
		killTimer := time.NewTimer(p.Lifetime)
		<-killTimer.C
		p.Kill()
		// spawnParitcle here
	}()

	return nil
}
