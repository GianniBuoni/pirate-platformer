package assets

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *AssetData) GetImage(key string) (rl.Texture2D, error) {
	image, ok := a.Images[key]
	if !ok {
		return rl.Texture2D{}, fmt.Errorf(
			"issue getting asset: %s. make sure it is loaded", key,
		)
	}
	return image, nil
}

func (a *AssetData) GetPlayer(key string) (rl.Texture2D, error) {
	image, ok := a.Player[key]
	if !ok {
		return rl.Texture2D{}, fmt.Errorf(
			"issue getting player: %s. make sure it is loaded", key,
		)
	}
	return image, nil
}
