package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		return UnmarshalError(path, err)
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
		return UnmarshalError(path, err)
	}
	obj := template.Objects
	a.SpawnIn[key] = obj
	return nil
}

func (a *Assets) importMapFile(key, path string) error {
	intKey64, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return err
	}
	intKey := int(intKey64)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	ld := LevelData{}
	err = json.Unmarshal(data, &ld)
	if err != nil {
		return UnmarshalError(path, err)
	}
	a.Maps[intKey] = ld
	return nil
}
