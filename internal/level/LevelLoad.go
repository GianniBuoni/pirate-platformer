package level

func (l *LevelData) Load() error {
	for _, objGroup := range l.mapData.ObjectGroups {
		if objGroup.Name == "zero index" {
			err := l.loadZero(objGroup.Objects)
			if err != nil {
				return err
			}
			err = l.loadTiles()
			if err != nil {
				return err
			}
		}
		if objGroup.Name == "static details" {
			err := l.loadStaticDetails(objGroup.Objects)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
