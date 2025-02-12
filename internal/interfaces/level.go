package interfaces

/*
Keeps track of level data: map, player, sprites...

Responisble for loading level data, drawing sprites,
tracking and resolving collisions, and updating event timers.
*/
type Level interface {
	Load() error
	Update()
	Draw() error
}
