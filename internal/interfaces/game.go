package interfaces

type Game interface {
  Load()
  Run()
  Quit()

  // getters
  IsRunning() bool

  // setters
  AddSprite(Sprite)
}
