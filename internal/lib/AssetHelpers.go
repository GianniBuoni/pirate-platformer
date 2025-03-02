package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func (a *Assets) importTileData(key, path string) error {
	// Inital import of raw data
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
	return nil
}

func (a *Assets) importSpawnIn(key, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	template := Template{}
	err = json.Unmarshal(data, &template)
	if err != nil {
		return fmt.Errorf("%s, %w", path, err)
	}
	obj := template.Objects
	a.SpawnIn[key] = obj
	return nil
}
