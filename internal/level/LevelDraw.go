package level

func (l *Level) Draw() error {
	groups := []string{"all", "ephemeral", "water"}
	for _, group := range groups {
		err := l.groups.Draw(group)
		if group == "water" && err != nil {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}
