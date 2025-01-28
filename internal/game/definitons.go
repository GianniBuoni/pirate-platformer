package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Running    bool
	player     *Player
	allSprites []Sprite

	PlayerTextures map[string]rl.Texture2D
	LevelTextures  map[string]rl.Texture2D
	UITextures     map[string]rl.Texture2D
}
