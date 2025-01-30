package assets

import (
	"github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetData struct {
	Images map[string]rl.Texture2D
	Frames map[string][]rl.Texture2D
}

func NewAssets() interfaces.Assets {
	return &AssetData{
		Images: map[string]rl.Texture2D{},
		Frames: map[string][]rl.Texture2D{},
	}
}

func (a *AssetData) Unload() {
	for _, image := range a.Images {
		rl.UnloadTexture(image)
	}

	for _, frames := range a.Frames {
		for _, image := range frames {
			rl.UnloadTexture(image)
		}
	}
}
