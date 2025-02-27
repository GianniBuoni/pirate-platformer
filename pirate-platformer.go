package main

import (
	"github.com/GianniBuoni/pirate-platformer/internal/game"
)

func main() {
	game := game.NewGame()
	game.Load()
	defer game.Quit(0)

	for game.Running {
		game.Run()
	}
}
