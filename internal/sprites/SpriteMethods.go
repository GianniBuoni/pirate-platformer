package sprites

import (
	"log"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *BasicSprite) Update() {}

func (s *BasicSprite) Draw(a Assets) {
	src, err := a.GetImage(ImageLib, s.image)
	if err != nil {
		log.Fatal(err)
	}

	rl.DrawTextureV(
		src,
		s.Pos(),
		rl.White,
	)
}
