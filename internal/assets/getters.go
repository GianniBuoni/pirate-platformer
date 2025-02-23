package assets

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) GetImage(
	aLib AssetLibrary, key string,
) (rl.Texture2D, error) {
	var (
		ok    bool
		image rl.Texture2D
	)
	switch aLib {
	case ImageLib:
		image, ok = a.Images[key]
	case PlayerLib:
		image, ok = a.Player[key]
	case TilesetLib:
		image, ok = a.Tilesets[key]
	}
	if !ok {
		return rl.Texture2D{}, fmt.Errorf(
			"error getting asset: '%s', make sure it is loaded", key,
		)
	}
	return image, nil
}
