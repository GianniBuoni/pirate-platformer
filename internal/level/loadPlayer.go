package level

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

func (l *LevelData) loadPlayer(obj *tiled.Object) error {
	newPlayer := sprites.NewPlayerParams{
		Assets: l.levelAssets,
		Pos:    rl.NewVector2(float32(obj.X), float32(obj.Y)),
		Groups: l.groups,
	}
	p, err := sprites.NewPlayer(newPlayer)
	if err != nil {
		return err
	}
	l.AddPlayer(p)
	return nil
}
