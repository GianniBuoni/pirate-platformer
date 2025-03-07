package level

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

func (l *Level) checkClouds() error {
	// 0 index is the water so we need to skip it
	ids, err := l.groups.GetIDs("clouds large")
	if err != nil {
		// an error here just means that the level might not have
		// large clouds
		return nil
	}
	ids = ids[1:]
	cloudsLarge, err := l.groups.GetSpritesID("clouds large", ids)
	if err != nil {
		return err
	}
	if cloudsLarge[0].HitBox().Right() <= 0 {
		for _, s := range cloudsLarge {
			x := s.HitBox().Left()
			w := s.HitBox().Width
			s.HitBox().Set(Left(x + w))
		}
	}
	return nil
}
