package level

import (
	"fmt"
	"os"
)

func (l *LevelData) Update() {
	for _, mSprite := range l.groups["moving"] {
		mSprite.Update()
	}
	l.player.Update()
	err := l.checkShells()
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(2)
	}
	l.cleanup("all", "moving", "damage")
}

func (l *LevelData) cleanup(groups ...string) {
	for _, group := range groups {
		for i, sprite := range l.groups[group] {
			if sprite.GetKill() {
				l.groups[group] = removeSprite(i, l.groups[group])
			}
		}
	}
}
