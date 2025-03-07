package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetLibrary uint

const (
	ImageLib AssetLibrary = iota
	PlayerLib
	TilesetLib
	UiLib
	TileData
	SpawnInLib
	FontLib
	MapLib
)

type Assets struct {
	// TEXTURES
	Images   map[string]rl.Texture2D
	Player   map[string]rl.Texture2D
	Tilesets map[string]rl.Texture2D
	UI       map[string]rl.Texture2D
	// DATA
	TilesetData map[string]Tileset
	SpawnIn     map[string]Object
	Hitboxes    map[string]Rect
	Fonts       map[string]rl.Font
	Maps        map[int]LevelData
}

func NewAssets() *Assets {
	return &Assets{
		Images:      map[string]rl.Texture2D{},
		Player:      map[string]rl.Texture2D{},
		Tilesets:    map[string]rl.Texture2D{},
		UI:          map[string]rl.Texture2D{},
		TilesetData: map[string]Tileset{},
		SpawnIn:     map[string]Object{},
		Hitboxes:    map[string]Rect{},
		Fonts:       map[string]rl.Font{},
		Maps:        map[int]LevelData{},
	}
}

func (a *Assets) Unload() {
	imagesToUnload := []map[string]rl.Texture2D{
		a.Images, a.Player, a.Tilesets, a.UI,
	}
	for _, lib := range imagesToUnload {
		for _, v := range lib {
			rl.UnloadTexture(v)
		}
	}
	for _, v := range a.Fonts {
		rl.UnloadFont(v)
	}
}
