package lib

import (
	"fmt"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
