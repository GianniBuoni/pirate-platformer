package level

import (
	"errors"
	"fmt"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func parseTile(idx, gid int, l *LevelData) (Tile, error) {
	t := Tile{}

	// parse image data
	var err error
	t.ImgX, t.ImgY, t.Image, err = parseTileGID(gid, l)
	if err != nil {
		return Tile{}, err
	}

	// parse tile position
	t.X = float32(idx%l.Width) * TileSize
	t.Y = float32(idx/l.Width) * TileSize

	return t, nil
}

func parseTileGID(gid int, l *LevelData) (
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
