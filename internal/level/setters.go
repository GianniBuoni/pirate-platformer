package level

import . "github.com/GianniBuoni/pirate-platformer/internal/interfaces"

func (l *LevelData) AddPlayer(s Sprite) {
	l.player = s
	l.AddSpriteGroup(s, "all")
}

func (l *LevelData) AddSpriteGroup(s Sprite, groups ...string) {
	for _, group := range groups {
		l.groups[group] = append(l.groups[group], s)
	}
}
