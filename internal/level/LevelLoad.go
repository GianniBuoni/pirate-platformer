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
		if objGroup.Name == "objects" {
			err := l.loadObjects(objGroup.Objects)
			if err != nil {
				return err
			}
		}
		if objGroup.Name == "moving" {
			err := l.loadMovingDetails(objGroup.Objects)
			if err != nil {
				return err
			}
			err = l.loadMoving(objGroup.Objects)
			if err != nil {
				return err
			}
		}
		if objGroup.Name == "player" {
			err := l.loadPlayer(objGroup.Objects[0])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
