package level

func (l *LevelData) Load() error {
	for _, objGroup := range l.mapData.ObjectGroups {
		if objGroup.Name == "zero index" {
			l.loadZero(objGroup.Objects)
		}
		/*
			if objGroup.Name == "BG details" {
				err := l.loadTiles(objGroup.Objects)
				if err != nil {
					return err
				}
				err = l.loadBGDetails(objGroup.Objects)
				if err != nil {
					return err
				}
			}
			if objGroup.Name == "Objects" {
				err := l.loadObjects(objGroup.Objects)
				if err != nil {
					return err
				}
			}
		*/
	}
	return nil
}
