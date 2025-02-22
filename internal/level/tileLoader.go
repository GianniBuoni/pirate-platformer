package level

import "fmt"

var bgTileLoader = Loader[[]int]{
	Run: loadBGTile,
}

var cTileLoader = Loader[[]int]{}

var pTileLoader = Loader[[]int]{}

func loadBGTile(data []int, l *LevelData) error {
	for _, id := range data {
		if id == 0 {
			continue
		}
		x, y, name, err := l.GetTileData(id)
		if err != nil {
			return err
		}
		fmt.Printf("name: %s, x: %v, y: %v\n", name, x, y)
	}
	return nil
}
