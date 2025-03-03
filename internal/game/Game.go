package game

import (
	"strconv"
	"strings"

	"github.com/GianniBuoni/pirate-platformer/internal/level"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/ui"
)

type GameData struct {
	assets *Assets
	window *WindowData
	level  *level.LevelData
	ui     *ui.UI
	stats  *Stats
	// Game States
	Running bool
	Paused  bool
}

func NewGame() *GameData {
	g := &GameData{
		assets:  NewAssets(),
		stats:   NewStats(),
		window:  NewWindow(),
		Running: true,
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
