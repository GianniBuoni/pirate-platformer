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

type Loader interface {
	Key() string
	Run(Object, AssetLibrary, GameModule) error
}

type TextLoader struct {
	key     string
	Builder func(Object, map[string]Text)
}

type SpriteLoader struct {
	key     string
	Builder func(Object, AssetLibrary, GameModule) ([]Sprite, error)
	Groups  []string
}

func (sl *SpriteLoader) Key() string {
	return sl.key
}

func (sl *SpriteLoader) Run(o Object, aLib AssetLibrary, gm GameModule) error {
	sprites, err := sl.Builder(o, aLib, gm)
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
