package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var bgTileLoader = Loader[[]int]{
	Run: tileLoaderMiddleware(loadBGTile),
}

var cTileLoader = Loader[[]int]{
	Run: tileLoaderMiddleware(loadCTile),
}

var pTileLoader = Loader[[]int]{
	Run: tileLoaderMiddleware(loadPTile),
}

func loadBGTile(s Sprite, l *LevelData) {
	l.AddSpriteGroup(s, "all")
}

func loadCTile(s Sprite, l *LevelData) {
	l.AddSpriteGroup(s, "all", "collision")
}

func loadPTile(s Sprite, l *LevelData) {
	l.AddSpriteGroup(s, "all", "platform")
}

func parseTile(idx, gid int, l *LevelData) (*Tile, error) {
	t := Tile{}

	// parse image data
	var err error
	t.ImgX, t.ImgY, t.Image, err = l.GetTileData(gid)
	if err != nil {
		return nil, err
	}

	// parse tile position
	t.X = float32(idx%l.Width) * TileSize
	t.Y = float32(idx/l.Width) * TileSize

	return &t, nil
}

func tileLoaderMiddleware(f func(Sprite, *LevelData)) func([]int, *LevelData) error {
	return func(data []int, l *LevelData) error {
		for idx, id := range data {
			if id == 0 {
				continue
			}
			t, err := parseTile(idx, id, l)
			if err != nil {
				return err
			}
			ts, err := sprites.NewTileSprite(*t, l.levelAssets)
			if err != nil {
				return err
			}
			f(ts, l)
		}
		return nil
	}
}
