package ui

import (
	"encoding/json"
	"os"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	updateGroup []Sprite
	sprites     map[string]Sprite
	assets      *Assets
	stats       *Stats
}

func NewUI(s *Stats, a *Assets) (*UI, error) {
	return &UI{
		updateGroup: []Sprite{},
		sprites:     map[string]Sprite{},
		assets:      a,
		stats:       s,
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
			if obj.Image != "" {
				s, err := sprites.NewSprite(obj, UiLib, ui.assets)
				if err != nil {
					return err
				}
				ui.sprites[obj.Image] = s
			}
		}
	}
	return nil
}

func (ui *UI) Update() error {
	if len(ui.updateGroup) > 0 {
		for _, s := range ui.updateGroup {
			err := s.Update()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ui *UI) Draw(gameUpdate bool) {
	if !gameUpdate {
		rl.DrawRectangleRec(
			rl.NewRectangle(0, 0, WindowW, WindowH),
			rl.ColorAlpha(WaterColor, 0.8),
		)
		ui.drawTextDisplay("GAME PAUSED", rl.NewVector2(WindowW/2, WindowH/2))
	}
	ui.drawStats()
}
