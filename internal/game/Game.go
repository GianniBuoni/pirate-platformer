package game

import (
	"strconv"

	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	"github.com/GianniBuoni/pirate-platformer/internal/level"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type GameData struct {
	levelAssets  *assets.Assets
	levelCurrent *level.LevelData
	levelMaps    map[int]string
	stats        *Stats
	window       *WindowData
	Running      bool
}

type Stats struct {
	currentLevel  int
	unlockedLevel int
	playerHealth  int
}

func NewGame() *GameData {
	return &GameData{
		levelMaps: GetMaps(),
		stats:     NewStats(),
		Running:   true,
	}
}

func NewStats() *Stats {
	return &Stats{
		currentLevel:  5,
		unlockedLevel: 4,
		playerHealth:  5,
	}
}

func GetMaps() map[int]string {
	levelMaps := map[int]string{}
	mapPaths := lib.GetFilePaths("data", "levels")
	for _, path := range mapPaths {
		strKey := lib.GetAssetKey(path)
		key, err := strconv.ParseInt(strKey, 10, 0)
		if err != nil {
			continue
		}
		levelMaps[int(key)] = path
	}
	return levelMaps
}
