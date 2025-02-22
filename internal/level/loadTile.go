package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var bgTileLoader = Loader[[]int]{
	Run: loadBGTile,
}

var cTileLoader = Loader[[]int]{}

var pTileLoader = Loader[[]int]{}

func loadBGTile(data []int, l *LevelData) error {
	for idx, id := range data {
		// parse image data
		if id == 0 {
			continue
		}
		t := Tile{}
		var err error
		t.ImgX, t.ImgY, t.Image, err = l.GetTileData(id)
		if err != nil {
			return err
		}

		// parse tile position
		t.X = float32(idx%l.Width) * TileSize
		t.Y = float32(idx/l.Width) * TileSize

		// init new tile sprite
		ts, err := sprites.NewTileSprite(t, l.levelAssets)
		if err != nil {
			return err
		}

		l.AddSpriteGroup(ts, "all")

	}
	return nil
}
