package assets

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *AssetData) ImportImages(
	aLib AssetLibrary,
	root ...string,
) error {
	paths := lib.GetFilePaths(root...)

	for _, path := range paths {
		key := strings.Split(filepath.Base(path), ".")[0]
		switch aLib {
		case ImageLib:
			a.Images[key] = rl.LoadTexture(path)
		case PlayerLib:
			a.Player[key] = rl.LoadTexture(path)
		case TilesetLib:
			a.Tilesets[key] = rl.LoadTexture(path)
		default:
			return errors.New("asset library not implemented.")
		}
	}
	return nil
}

func (a *AssetData) GetImage(
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
