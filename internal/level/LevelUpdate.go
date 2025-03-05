package level

import "fmt"

func (l *Level) Update() error {
	if l.stats.Paused {
		return nil
	}

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

	err := l.player.Update()
	if err != nil {
		return err
	}
	l.camera.Update()

	// manage sprites
	err = l.spawnInOut()
	if err != nil {
		return err
	}
	return nil
}
