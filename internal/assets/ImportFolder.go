package assets

import (
	"path/filepath"
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *AssetData) ImportFolder(
	rootArgs []string,
) {

	paths := lib.GetFilePaths(rootArgs)

	for _, path := range paths {
		key := strings.Split(filepath.Base(path), ".png")[0]
		a.Images[key] = rl.LoadTexture(path)
	}
}
