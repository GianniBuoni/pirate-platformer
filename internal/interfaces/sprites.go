package interfaces

type Sprites interface {
	Update()
	Draw(Assets) error
}
