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
	return nil
}
