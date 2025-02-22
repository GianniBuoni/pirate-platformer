package level

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

var bgTileLoader = Loader[[]int]{
	Key: "bg",
	Run: tileLoaderMiddleware(loadBGTile),
}

var cTileLoader = Loader[[]int]{
	Key: "collision",
	Run: tileLoaderMiddleware(loadCTile),
}

var pTileLoader = Loader[[]int]{
	Key: "platform",
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
			ts, err := sprites.NewTileSprite(t, l.levelAssets)
			if err != nil {
				return err
			}
			f(ts, l)
		}
		return nil
	}
}
