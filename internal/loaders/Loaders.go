package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type GameModule interface {
	Assets() *Assets
	Texts() map[string]Text
	Sprites() map[int]Sprite
	NextId() int
	AddSpriteGroup(Sprite, map[int]Sprite, ...string)
}

type Loader[T any] interface {
	Key() string
	Run(T, AssetLibrary, GameModule) error
}

type TextLoader struct {
	key     string
	Builder func(Object, map[string]Text)
}

type SpriteLoader[T any] struct {
	key     string
	Builder func(T, AssetLibrary, GameModule) ([]Sprite, error)
	Groups  []string
}

func (sl *SpriteLoader[T]) Key() string {
	return sl.key
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
	return nil
}

func (tl *TextLoader) Key() string {
	return tl.key
}

func (tl *TextLoader) Run(o Object, aLib AssetLibrary, gm GameModule) error {
	tl.Builder(o, gm.Texts())
	return nil
}
