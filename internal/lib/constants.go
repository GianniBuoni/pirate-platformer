package lib

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// GAME
	WindowW  float32 = 1280
	WindowH  float32 = 720
	TileSize float32 = 64
	Title            = "Pirate Platformer 💀!"

	// BASIC SPRITES
	PlatHitbox float32 = 13

	// ANIMATED SPRITES
	FrameSpeed float32 = 6

	// PLAYER
	Gravity     float32 = 1200
	PlayerSpeed float32 = 300
	JumpDist    float32 = 700
)

var (
	BgColor rl.Color = rl.NewColor(221, 198, 161, 255)
)
