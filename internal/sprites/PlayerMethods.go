package sprites

import (
	"log"

	"github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) Update() {
	p.animate()
}

func (p *Player) Draw(a interfaces.Assets) {
	src, err := a.GetImage(interfaces.Player, p.image)
	if err != nil {
		log.Fatal(err)
	}

	srcRect := rl.NewRectangle(
		p.frameOffset*float32(int(p.frameIndex)%p.frameCount),
		p.frameOffset*float32(int(p.frameIndex)%p.frameCount),
		p.frameOffset,
		p.frameOffset,
	)

	rl.DrawTexturePro(
		src,
		srcRect,
		p.rect,
		rl.Vector2Scale(p.posOffset, 2),
		0,
		rl.White,
	)
}

func (p *Player) animate() {
	dt := rl.GetFrameTime()
	p.frameIndex += p.frameSpeed * dt
}
