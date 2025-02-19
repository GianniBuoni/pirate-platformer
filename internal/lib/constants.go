package lib

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WindowW     float32 = 1280
	WindowH     float32 = 720
	TileSize    float32 = 64
	Gravity     float32 = 1200
	PlayerSpeed float32 = 300
	JumpDist    float32 = 700
	PlatHitbox  float32 = 13
	FrameSpeed          = 6
	Title               = "Pirate Platformer 💀!"
)

var (
	BgColor rl.Color = rl.NewColor(221, 198, 161, 255)
)
