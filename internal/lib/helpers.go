package lib

import (
	"errors"
	"io/fs"
	"log"
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
				log.Fatalf("issue walking through dir: %v", err)
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
			"error generating random range: max < min, or max == min",
		)
	}
	return rand.Intn(max-min) + min, nil
}
