package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) ImportImages(
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

func (a *Assets) ImportTilesetData(root ...string) error {
	paths := lib.GetFilePaths(root...)
	for _, path := range paths {
		if !strings.Contains(path, "json") {
			continue
		}
		key := lib.GetAssetKey(path)

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		tsd := lib.Tileset{}
		err = json.Unmarshal(data, &tsd)
		if err != nil {
			return fmt.Errorf("%s, %w", path, err)
		}
		a.TilesetData[key] = tsd
	}
	return nil
}
