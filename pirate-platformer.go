package main

import (
	"github.com/GianniBuoni/pirate-platformer/internal/game"
)

func main() {
	game := game.NewGame()
	game.LoadLevel()
	defer game.Quit(0)

	for game.Running {
		game.Run()
	}
}
