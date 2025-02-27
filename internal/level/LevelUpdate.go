package level

func (l *LevelData) Update() error {
	for _, mSprite := range l.groups["moving"] {
		mSprite.Update()
	}
	l.player.Update()
	err := l.checkShells()
	if err != nil {
		return err
	}
	l.cleanup("all", "moving", "damage")
	return nil
}

func (l *LevelData) cleanup(groups ...string) {
	for _, group := range groups {
		for i, sprite := range l.groups[group] {
			if sprite.GetKill() {
				l.groups[group] = removeSprite(i, l.groups[group])
			}
		}
	}
}
