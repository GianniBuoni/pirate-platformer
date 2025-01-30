package assets

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *AssetData) GetImage(key string) (rl.Texture2D, error) {
	image, ok := a.Images[key]
	if !ok {
		return rl.Texture2D{}, errors.New("issue getting asset. make sure it is loaded")
	}
	return image, nil
}
