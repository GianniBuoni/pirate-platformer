package level

import "fmt"

func (l *Level) Draw() {
	fmt.Printf("l.groups[\"all\"]: %v\n", l.groups["all"])
	for _, id := range l.groups["all"] {
		s := l.spirtes[id]
		s.Draw(s.GetID().Src, s.GetPos())
	}
}
