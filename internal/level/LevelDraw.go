package level

func (l *Level) Draw() {
	for _, id := range l.groups["all"] {
		s := l.spirtes[id]
		s.Draw(s.GetID().Src, s.GetPos())
	}
}
