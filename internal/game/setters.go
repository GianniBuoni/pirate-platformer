package game

import "github.com/GianniBuoni/pirate-platformer/internal/interfaces"

func (g *GameData) AddSprite(s interfaces.Sprite) {
	g.allSprites = append(g.allSprites, s)
}

func (g *GameData) AddPlayer(s interfaces.Sprite) {
	g.player = s
	g.AddSprite(s)
}
