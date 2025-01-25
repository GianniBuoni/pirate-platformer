package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (a *Assets) ImportFolder(pathArgs []string) error {
	rootPath := filepath.Join(pathArgs...)
	a.mu.Lock()
	defer a.mu.Unlock()

	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("issue parsing folder path: %w", err)
		}

		frames := []rl.Texture2D{}
		dirName := pathArgs[len(pathArgs)-1]

		if !info.IsDir() {
			frames = append(frames, rl.LoadTexture(path))
		}

		a.Frames[dirName] = frames

		return nil
	})

	return nil
}
