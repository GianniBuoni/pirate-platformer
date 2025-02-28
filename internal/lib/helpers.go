package lib

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"path/filepath"
	"strings"
)

func GetFilePaths(args ...string) (paths []string) {
	root := filepath.Join(args...)

	filepath.Walk(
		root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("Issue walking through path: %w", err)
			}

			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		},
	)
	return paths
}

func GetAssetKey(fullPath string) string {
	return strings.Split(filepath.Base(fullPath), ".")[0]
}

func RandInt(min, max int) (n int, err error) {
	if max < min || max == min {
		return 0, errors.New(
			"Invalid random number range: max < min, or max == min",
		)
	}
	return rand.Intn(max-min) + min, nil
}
