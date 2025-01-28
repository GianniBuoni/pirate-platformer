package game

import (
	"github.com/GianniBuoni/pirate-platformer/internal/player"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Running    bool //game state
	Player     *player.Player
	AllSprites []sprites.Sprite

	PlayerTextures map[string]rl.Texture2D //assets
	LevelTextures  map[string]rl.Texture2D
	UITextures     map[string]rl.Texture2D
}

func NewGame() *Game {
	return &Game{
		Running:        true,
		PlayerTextures: map[string]rl.Texture2D{},
		LevelTextures:  map[string]rl.Texture2D{},
	}
}

func (g *Game) Load() {
	// TODO: preload index of assets to game
}

func (g *Game) Unload() {
	for _, v := range g.PlayerTextures {
		rl.UnloadTexture(v)
	}
	for _, v := range g.LevelTextures {
		rl.UnloadTexture(v)
	}
}

func (g *Game) Update() {
	for _, sprite := range g.AllSprites {
		sprite.Update()
	}
}

func (g *Game) Draw() {
	for _, sprite := range g.AllSprites {
		sprite.Draw()
	}
}
