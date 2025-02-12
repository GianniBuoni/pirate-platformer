package interfaces

import rl "github.com/gen2brain/raylib-go/raylib"

type AssetLibrary uint

const (
	Images AssetLibrary = iota
	Player
	Tilesets
	Frames
)

type Assets interface {
	// Import contents of folder into a specified asset library.
	// Walks through the provided folder and proccesses all files,
	// and creates a key from the file name
	ImportImages(AssetLibrary, ...string) error
	// Returns a raylib texture from the asset library.
	// Takes a key (file name) and an AssetLibrary.
	GetImage(AssetLibrary, string) (rl.Texture2D, error)
	// Frees assets from memory
	Unload()
}
