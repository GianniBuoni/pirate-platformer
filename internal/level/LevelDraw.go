package level

import "fmt"

func (l *Level) Draw() {
	for _, id := range l.groups["all"] {
		s := l.spirtes[id]
		s.Draw(s.GetID().Src, s.GetPos())
	}

	ephemeral, ok := l.groups["ephemeral"]
	if ok {
		for _, id := range ephemeral {
			s, ok := l.spirtes[id]
			if !ok {
				fmt.Printf("Can't draw %d; it might be deleted\n", id)
				continue
			}
			s.Draw(s.GetID().Src, s.GetPos())
		}
	}
}
