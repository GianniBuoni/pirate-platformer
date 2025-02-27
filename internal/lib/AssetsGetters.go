package lib

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

func (a *Assets) GetObject(key string) (Object, error) {
	obj, ok := a.SpawnIn[key]
	if !ok {
		return Object{},
			fmt.Errorf(
				"spawn in asset key: %s not found, make sure it's loaded\n",
				key,
			)
	}
	return obj, nil
}

func (a *Assets) GetTileset(key string) (Tileset, error) {
	tileset, ok := a.TilesetData[key]
	if !ok {
		return Tileset{},
			fmt.Errorf("tileset key: %s not found, make sure it's loaded\n", key)
	}
	return tileset, nil
}

func (a *Assets) GetHitbox(key string) (Rect, bool) {
	hitbox, ok := a.Hitboxes[key]
	if !ok {
		return Rect{}, false
	}
	return hitbox, true
}
