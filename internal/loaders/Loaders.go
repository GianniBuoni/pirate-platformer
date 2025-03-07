package loaders

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type GameModule interface {
	Assets() *Assets
	Texts() map[string]Text
	AddSpriteGroup(Sprite, ...string)
}

type Loader[T any] interface {
	GetKey() string
	Run(T, AssetLibrary, GameModule) error
}

type TextLoader struct {
	Key     string
	Builder func(Object, map[string]Text)
}

type SpriteLoader[T any] struct {
	Key     string
	Builder func(T, AssetLibrary, GameModule) ([]Sprite, error)
	Groups  []string
}

func (sl *SpriteLoader[T]) GetKey() string {
	return sl.Key
}

func (sl *SpriteLoader[T]) Run(t T, aLib AssetLibrary, gm GameModule) error {
	sprites, err := sl.Builder(t, aLib, gm)
	if err != nil {
		return err
	}
	if sprites != nil {
		for _, s := range sprites {
			gm.AddSpriteGroup(s, sl.Groups...)
		}
	}
	return nil
}

func (tl *TextLoader) GetKey() string {
	return tl.Key
}

func (tl *TextLoader) Run(o Object, aLib AssetLibrary, gm GameModule) error {
	tl.Builder(o, gm.Texts())
	return nil
}
