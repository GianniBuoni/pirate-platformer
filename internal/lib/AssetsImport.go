package lib

import (
	"errors"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) ImportImages(
	aLib AssetLibrary,
	root ...string,
) error {
	paths := GetFilePaths(root...)

	for _, path := range paths {
		key := GetAssetKey(path)
		switch aLib {
		case ImageLib:
			a.Images[key] = rl.LoadTexture(path)
		case PlayerLib:
			a.Player[key] = rl.LoadTexture(path)
		case TilesetLib:
			a.Tilesets[key] = rl.LoadTexture(path)
		case UiLib:
			a.UI[key] = rl.LoadTexture(path)
		default:
			return errors.New("passed in asset library not implemented.")
		}
	}
	return nil
}

func (a *Assets) ImportData(aLib AssetLibrary, root ...string) error {
	paths := GetFilePaths(root...)
	for _, path := range paths {
		if !strings.Contains(path, "json") && !strings.Contains(path, "ttf") {
			continue
		}
		key := GetAssetKey(path)
		switch aLib {
		case TileData:
			err := a.importTileData(key, path)
			if err != nil {
				return err
			}
		case SpawnInLib:
			err := a.importSpawnIn(key, path)
			if err != nil {
				return err
			}
		case MapLib:
			err := a.importMapFile(key, path)
			if err != nil {
				return err
			}
		case FontLib:
			a.Fonts[key] = rl.LoadFont(path)
		default:
			return errors.New("passed in asset library not implemented.")
		}
	}
	return nil
}
