package lib

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// GAME
	WindowW  float32 = 1280
	WindowH  float32 = 720
	TileSize float32 = 64
	Title            = "Pirate Platformer 💀!"

	// UI
	UITileSize float32 = 32

	// ANIMATED SPRITES
	FrameSpeed float32 = 6

	// PLAYER
	MaxHealth int     = 10
	Gravity   float32 = 2
)

var (
	BgColor    rl.Color = rl.NewColor(221, 198, 161, 255)
	WaterColor rl.Color = rl.NewColor(146, 169, 206, 255)
)
