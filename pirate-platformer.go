package main

import (
	"github.com/GianniBuoni/pirate-platformer/internal/game"
)

func main() {
	game := game.NewGame()
	game.Load()
	defer game.Quit()

	for game.Running {
		game.Run()
	}
}
