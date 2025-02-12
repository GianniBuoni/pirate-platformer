package level

import . "github.com/GianniBuoni/pirate-platformer/internal/interfaces"

func (l *LevelData) AddSprite(s Sprite) {
	l.allSprites = append(l.allSprites, s)
}

func (l *LevelData) AddPlayer(s Sprite) {
	l.player = s
	l.AddSprite(s)
}
