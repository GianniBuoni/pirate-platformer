package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Item struct {
	Pos
	ID
	Animation
	Value int
}

func NewItem(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	item := &Item{
		Pos:       newPos(obj, a),
		Animation: newAnimation(),
		Value:     obj.Properties.Value,
	}
	var err error
	item.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Item) Update() error {
	item.animate(item.rect, item.Src)
	return nil
}
