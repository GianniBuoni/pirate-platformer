package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *LevelData) Update() error {
	for _, mSprite := range l.groups["moving"] {
		err := mSprite.Update()
		if err != nil {
			return err
		}
	}
	err := l.player.Update()
	if err != nil {
		return err
	}
	l.camera.Update()
	err = l.checkShells()
	if err != nil {
		return err
	}
	l.itemCollisions()
	// TODO: make unloader for each loader?
	// use id to remove from groups?
	// rework groups to create faster lookups?
	l.checkPearls()
	l.checkClouds()
	l.cleanup("all", "moving", "damage", "pearl")
	return nil
}

func (l *LevelData) checkShells() error {
	for _, s := range l.groups["shell"] {
		shell, ok := s.(*Shell)
		if !ok {
			return fmt.Errorf(
				"Sprite %s, in sprite group \"shell\" is not a shell sprite",
				s.GetID().Image,
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

func (l *LevelData) checkPearls() {
	for _, p := range l.groups["pearl"] {
		if p.GetID().Kill {
			l.spawnParticle(p)
		}
	}
}

func (l *LevelData) checkClouds() {
	_, ok := l.groups["clouds large"]
	if ok && l.groups["clouds large"][1].Rect().Right() <= 0 {
		for i, s := range l.groups["clouds large"] {
			if i > 0 {
				s.Rect().X += s.Rect().Width
			}
		}
	}
}

func (l *LevelData) cleanup(groups ...string) {
	for _, group := range groups {
		for i, sprite := range l.groups[group] {
			if sprite.GetID().Kill {
				l.groups[group] = removeSprite(i, l.groups[group])
			}
		}
	}
}
