package game

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
)

type GameData struct {
	running      bool
	window       Window
	levelAssets  Assets
	player       Sprite
	allSprites   []Sprite
}

func NewGame() Game {
	return &GameData{
		running: true,
	}
}
