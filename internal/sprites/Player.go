package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Player struct {
	ID
	Pos
}

func NewPlayer(obj Object, a *assets.Assets) (*Player, error) {
	id, err := newId("fall", assets.PlayerLib, a)
	if err != nil {
		return nil, err
	}
	p := Player{
		ID:  id,
		Pos: newPos(obj.X, obj.Y, obj.Width, obj.Height),
	}
	return &p, nil
}
