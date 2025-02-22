package level

import (
	"errors"
	"fmt"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (l *LevelData) GetTileData(gid int) (
	x, y float32, name string, err error,
) {
	var firstGID int

	for k, v := range l.tileRefs {
		if k.FirstGID <= gid && k.LastGID >= gid {
			name = v
			firstGID = k.FirstGID
		}
	}
	if name == "" {
		return 0, 0, "", fmt.Errorf("gid: %d\n not found in tileset refs.", gid)
	}
	tileset := l.levelAssets.TilesetData[name]

	idx := gid - firstGID
	if idx > tileset.Count {
		fmt.Printf("index: %d out of range of tileset '%s'.\n", idx, name)
		return 0, 0, "",
			errors.New("Check if there are rotation flags on tile gid")
	}
	x = float32(idx%tileset.Columns) * TileSize
	y = float32(idx/tileset.Columns) * TileSize
	return x, y, name, nil
}
