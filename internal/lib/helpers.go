package lib

import (
	"io/fs"
	"log"
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
