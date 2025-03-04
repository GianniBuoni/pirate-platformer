package sprites

import (
	"strconv"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func NewPath(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	s := ObjectSprite{
		ID: ID{
			Image: obj.Image,
		},
		Pos: newPos(obj, a),
	}
	return &s, nil
}

func (s *ObjectSprite) SetPaths(dest map[int]*Rect) error {
	key, err := strconv.ParseInt(s.Image, 10, 64)
	if err != nil {
		return err
	}
	dest[int(key)] = s.Rect()
	return nil
}
