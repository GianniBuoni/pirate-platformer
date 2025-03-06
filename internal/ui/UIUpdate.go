package ui

func (ui *UI) Update() error {
	// update ephemeral group
	err := ui.groups.Update("ephemeral")
	if err != nil {
		return err
	}
	// update heart group
	err = ui.updateHearts()
	if err != nil {
		return err
	}
	// cleanup
	err = ui.groups.Cleanup("heart")
	if err != nil {
		return err
	}
	err = ui.groups.Cleanup("ephemeral")
	if err != nil {
		return err
	}
	return nil
}
