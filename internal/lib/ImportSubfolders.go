package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) ImportSubfolders(pathArgs []string) error {
	rootPath := filepath.Join(pathArgs...)
	a.mu.Lock()
	defer a.mu.Unlock()

	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("issue parsing folder path: %w", err)
		}

		if !info.IsDir() {
			key := fmt.Sprintf(
				"%s_%s",
				pathArgs[len(pathArgs)-1],
				filepath.Base(filepath.Dir(path)),
			)

			mapFrames, ok := a.Frames[key]
			if !ok {
				a.Frames[key] = []rl.Texture2D{rl.LoadTexture(path)}
			} else {
				mapFrames = append(mapFrames, rl.LoadTexture(path))
			}
		}
		return nil
	})
	return nil
}
