package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Item struct {
	Pos
	ID
	Animation
	Value int
}

func NewItem(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	id, err := newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return &Item{
		Pos:       newPos(obj, a),
		ID:        id,
		Animation: newAnimation(),
		Value:     obj.Properties.Value,
	}, nil
}
