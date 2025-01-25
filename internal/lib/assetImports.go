package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Assets struct {
	mu     sync.RWMutex
	Frames map[string][]rl.Texture2D
}

type ImportImageOptions struct {
	Format string
}

func (a *Assets) ImportImage(
	path []string,
	opts ...func(*ImportImageOptions),
) error {

	iio := &ImportImageOptions{
		Format: "png",
	}

	for _, opt := range opts {
		opt(iio)
	}

	fullPath := filepath.Join(path...) + fmt.Sprintf(".%s", iio.Format)
	fileName := path[len(path)-1]

	a.mu.Lock()
	defer a.mu.Unlock()

	frames := []rl.Texture2D{rl.LoadTexture(fullPath)}
	fmt.Println(fileName, frames)
	a.Frames[fileName] = frames
	return nil
}

func WithFormat(format string) func(*ImportImageOptions) {
	return func(iio *ImportImageOptions) {
		iio.Format = format
	}
}

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
