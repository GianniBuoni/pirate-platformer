package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) Update() {
	p.animate()
}

func (p *Player) Draw(a Assets) error {
	src, err := a.GetImage(PlayerLib, p.image)
	if err != nil {
		return err
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
	return nil
}

func (p *Player) animate() {
	dt := rl.GetFrameTime()
	p.frameIndex += p.frameSpeed * dt
}
