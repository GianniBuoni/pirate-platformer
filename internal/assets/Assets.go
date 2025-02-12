package assets

import (
	"github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetData struct {
	Images   map[string]rl.Texture2D
	Player   map[string]rl.Texture2D
	Tilesets map[string]rl.Texture2D
}

func NewAssets() interfaces.Assets {
	return &AssetData{
		Images:   map[string]rl.Texture2D{},
		Player:   map[string]rl.Texture2D{},
		Tilesets: map[string]rl.Texture2D{},
	}
}

func (a *AssetData) Unload() {
	for _, image := range a.Images {
		rl.UnloadTexture(image)
	}
	for _, state := range a.Player {
		rl.UnloadTexture(state)
	}
	for _, tileset := range a.Tilesets {
		rl.UnloadTexture(tileset)
	}
}
