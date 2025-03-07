package level

func (l *Level) Update() (err error) {
	if l.stats.Paused {
		return nil
	}
	groups := []string{"moving", "ephemeral"}
	for _, group := range groups {
		err := l.groups.Update(group)
		if err != nil {
			return err
		}
	}
	// update player
	err = l.player.Update()
	if err != nil {
		return err
	}
	// update camera based on player update
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
	// manage sprite spawning
	err = l.spawnInOut()
	if err != nil {
		return err
	}
	// cleanup
	err = l.groups.Cleanup("ephemeral", "damage", "item")
	return nil
}
