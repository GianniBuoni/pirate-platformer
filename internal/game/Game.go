package game

import (
	"strconv"
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/ui"
)

type GameData struct {
	levelAssets  *Assets
	levelCurrent *level.LevelData
	levelMaps    map[int]string
	loaders      *level.Loaders
	ui           *ui.UI
	stats        *Stats
	window       *WindowData
	Running      bool
	CanUpdate    bool
}

func NewGame() *GameData {
	g := &GameData{
		levelAssets: NewAssets(),
		levelMaps:   GetMaps(),
		loaders:     level.NewLoaders(),
		stats:       NewStats(),
		window:      NewWindow(),
		Running:     true,
		CanUpdate:   true,
	}
	g.loadAssets()
	return g
}

func GetMaps() map[int]string {
	levelMaps := map[int]string{}
	mapPaths := GetFilePaths("data", "levels")
	for _, path := range mapPaths {
		if !strings.Contains(path, "json") {
			continue
		}
		strKey := GetAssetKey(path)
		key, err := strconv.ParseInt(strKey, 10, 0)
		if err != nil {
			continue
		}
		levelMaps[int(key)] = path
	}
	return levelMaps
}
