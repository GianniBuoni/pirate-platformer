package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetLibrary uint

const (
	ImageLib AssetLibrary = iota
	PlayerLib
	TilesetLib
)

type Assets struct {
	Images      map[string]rl.Texture2D
	Player      map[string]rl.Texture2D
	Tilesets    map[string]rl.Texture2D
	TilesetData map[string]Tileset
	Hitboxes    map[string]Rect
}

func NewAssets() *Assets {
	return &Assets{
		Images:      map[string]rl.Texture2D{},
		Player:      map[string]rl.Texture2D{},
		Tilesets:    map[string]rl.Texture2D{},
		TilesetData: map[string]Tileset{},
		Hitboxes:    map[string]Rect{},
	}
}

func (a *Assets) Unload() {
	for _, image := range a.Images {
		rl.UnloadTexture(image)
	}
	for _, state := range a.Player {
		rl.UnloadTexture(state)
	}
	for _, tileset := range a.Tilesets {
		rl.UnloadTexture(tileset)
	}
}
