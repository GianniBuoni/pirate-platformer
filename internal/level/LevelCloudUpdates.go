package level

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

func (l *Level) checkClouds() error {
	// 0 index is the water so we need to skip it
	firstID := l.groups["clouds large"][1]
	firstCloud, ok := l.spirtes[firstID]
	if !ok {
		return nil
	}
	if firstCloud.HitBox().Right() <= 0 {
		for i, id := range l.groups["clouds large"] {
			// 0 index is the water so we need to skip it
			if i == 0 {
				continue
			}
			s, ok := l.spirtes[id]
			if !ok {
				return fmt.Errorf(
					"large cloud spirte id \"%d\" not found in level sprite map", id,
				)
			}
			x := s.HitBox().Left()
			w := s.HitBox().Width
			s.HitBox().Set(Left(x + w))
		}
	}
	return nil
}
