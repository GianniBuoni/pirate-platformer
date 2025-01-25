package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Assets struct {
	assets map[string][]rl.Texture2D
	mu     sync.RWMutex
}

func importImage(a *Assets, path ...string) error {
	fullPath := filepath.Join(path...)
	fileName := strings.Split(filepath.Base(fullPath), ".")[0]

	a.mu.Lock()
	defer a.mu.Unlock()

	frames := []rl.Texture2D{rl.LoadTexture(fullPath)}
	a.assets[fileName] = frames

	return nil
}

func importFolder(a *Assets, pathArgs ...string) error {
	rootPath := filepath.Join(pathArgs...)
	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		a.mu.Lock()
		defer a.mu.Unlock()

		if err != nil {
			return fmt.Errorf("issue parsing folder path: %w", err)
		}

		frames := []rl.Texture2D{}
		dirName := pathArgs[len(pathArgs)-1]

		if !info.IsDir() {
			frames = append(frames, rl.LoadTexture(path))
		}

		a.assets[dirName] = frames

		return nil
	})
	return nil
}
