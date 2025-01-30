package interfaces

import rl "github.com/gen2brain/raylib-go/raylib"

type Assets interface {
	Unload()

	// setters
	ImportImages([]string)
	ImportPlayer([]string)

	// getters
	GetImage(string) (rl.Texture2D, error)
	GetPlayer(string) (rl.Texture2D, error)
}
