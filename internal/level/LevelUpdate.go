package level

import "fmt"

func (l *Level) Update() error {
	if l.stats.Paused {
		return nil
	}
	// update moving sprites
	moving, ok := l.groups["moving"]
	if !ok {
		fmt.Println("Level sprite group \"moving\" is empty.")
		return nil
	}
	for _, id := range moving {
		err := l.spirtes[id].Update()
		if err != nil {
			return err
		}
	}
	// update ephemeral sprites
	ephemeral, ok := l.groups["ephemeral"]
	if ok {
		for _, id := range ephemeral {
			s, ok := l.spirtes[id]
			if !ok {
				fmt.Printf("Can't update %d; it might be deleted\n", id)
				continue
			}
			err := s.Update()
			if err != nil {
				return err
			}
		}
	}
	// update player
	err := l.player.Update()
	if err != nil {
		return err
	}
	l.camera.Update()
	// check collisions
	err = l.itemCollisions()
	if err != nil {
		return err
	}
	// reset pos of large clouds if needed
	err = l.checkClouds()
	if err != nil {
		return err
	}
	// manage sprites
	err = l.spawnInOut()
	if err != nil {
		return err
	}
	return nil
}
