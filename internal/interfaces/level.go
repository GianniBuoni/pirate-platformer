package interfaces

import rl "github.com/gen2brain/raylib-go/raylib"

/*
Keeps track of level data: map, player, sprites...

Responisble for loading level data, drawing sprites,
tracking and resolving collisions, and updating event timers.
*/
type Level interface {
	Load() error
	Update()
	Draw() error
	// Reports player position (rl.Vector2) back to the main game
	PlayerPos() rl.Vector2
}
