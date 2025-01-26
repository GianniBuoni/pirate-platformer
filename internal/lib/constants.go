package lib

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// game & level info
	WindowWidth    int32 = 1280
	WindowHeight   int32 = 720
	TileSize       int32 = 64
	AnimationSpeed int32 = 6
)

var (
	// colors
	BgColor = rl.NewColor(221, 198, 161, 255)

	// layer sorting indexes
	ZLayers = map[string]int{
		"bg":         0,
		"clouds":     1,
		"bg tiles":   2,
		"path":       3,
		"bg details": 4,
		"main":       5,
		"water":      6,
		"fg":         7,
	}

	ItemValues = map[string]int{
		"potion":  1,
		"silver":  1,
		"gold":    5,
		"diamond": 20,
		"skull":   50,
	}
)
