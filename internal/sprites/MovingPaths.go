package sprites

import (
	"strconv"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func NewPath(obj Object, a *Assets) (Sprite, error) {
	s := ObjectSprite{
		ID: ID{
			image: obj.Image,
		},
		Pos: newPos(obj, a),
	}
	return &s, nil
}

func (s *ObjectSprite) SetPaths(dest map[int]*Rect) error {
	key, err := strconv.ParseInt(s.image, 10, 64)
	if err != nil {
		return err
	}
	dest[int(key)] = s.Rect()
	return nil
}
