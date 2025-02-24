package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) ImportImages(
	aLib AssetLibrary,
	root ...string,
) error {
	paths := GetFilePaths(root...)

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

func (a *Assets) ImportTilesetData(root ...string) error {
	// Inital import of raw data
	paths := GetFilePaths(root...)
	for _, path := range paths {
		if !strings.Contains(path, "json") {
			continue
		}
		key := GetAssetKey(path)

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		tsd := Tileset{}
		err = json.Unmarshal(data, &tsd)
		if err != nil {
			return fmt.Errorf("%s, %w", path, err)
		}
		a.TilesetData[key] = tsd

		// Further parsing of tile data into map of hitboxes
		if len(tsd.Tiles) > 0 {
			for _, tile := range tsd.Tiles {
				if len(tile.ObjectGroup.Hitboxes) > 0 {
					key := GetAssetKey(tile.Image)
					if len(tile.ObjectGroup.Hitboxes) > 1 {
						return fmt.Errorf("'%s' has more than one hitbox defined!", key)
					}
					hitbox := tile.ObjectGroup.Hitboxes[0]
					rect := NewRectangle(hitbox.X, hitbox.Y, hitbox.Width, hitbox.Height)
					a.Hitboxes[key] = *rect
				}
			}
		}
	}
	return nil
}
