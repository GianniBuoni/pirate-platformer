package ui

import (
	"encoding/json"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type UI struct {
	sprites map[string]Sprite
	assets  *Assets
	stats   *Stats
}

func NewUI(s *Stats, a *Assets) (*UI, error) {
	return &UI{
		sprites: map[string]Sprite{},
		assets:  a,
		stats:   s,
	}, nil
}

func (ui *UI) Load(mapPath string) error {
	// find get map objects and make sprites
	data, err := os.ReadFile(mapPath)
	if err != nil {
		return err
	}
	layerData := Layers{}
	err = json.Unmarshal(data, &layerData)
	if err != nil {
		return err
	}
	for _, layer := range layerData.Layers {
		if len(layer.Objects) == 0 {
			continue
		}
		for _, obj := range layer.Objects {
			s, err := sprites.NewSprite(obj, UiLib, ui.assets)
			if err != nil {
				return err
			}
			ui.sprites[obj.Image] = s
		}
	}
	return nil
}

func (ui *UI) Draw() {
	s := ui.sprites["heart"]
	c := ui.sprites["coin"]
	c.Draw(c.GetID(), c.GetPos())

	for i := range ui.stats.PlayerHP() {
		var margin float32
		if i == 0 {
			margin = 32
		} else {
			margin = 32 + (float32(i) * 8)
		}
		pos := s.GetPos()
		pos.Rect().X = float32(i)*pos.Rect().Width + margin
		s.Draw(s.GetID(), s.GetPos())
	}
}
