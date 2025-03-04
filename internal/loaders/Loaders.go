package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type GameModule interface {
	Load(string) error
	Update() error
	Draw()
	Assets() *Assets
	Sprites() map[int]Sprite
	AddSpriteGroup(Sprite, map[int]Sprite, ...string)
}

type SpriteLoader[T any] struct {
	Key     string
	Builder func(T, AssetLibrary, GameModule) ([]Sprite, error)
	Groups  []string
}

func (sl *SpriteLoader[T]) Run(t T, aLib AssetLibrary, gm GameModule) error {
	sprites, err := sl.Builder(t, aLib, gm)
	if err != nil {
		return err
	}
	if sprites != nil {
		for _, s := range sprites {
			gm.AddSpriteGroup(s, gm.Sprites(), sl.Groups...)
		}
	}
	return err
}
