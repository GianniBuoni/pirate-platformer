package interfaces

type Sprite interface {
	Update()
	Draw(Assets)

	// setters
	OffsetCentre()
}
