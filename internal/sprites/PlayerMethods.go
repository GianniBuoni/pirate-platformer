package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	"github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) Update() {
	p.input()
	p.move()
	p.animate()
	p.rect.X = p.hitbox.X - lib.TileSize
	p.rect.Y = p.hitbox.Y - lib.TileSize
	p.oldRect = p.hitbox
}

func (p *PlayerData) Draw(a Assets) error {
	src, err := a.GetImage(PlayerLib, p.image)
	if err != nil {
		return err
	}

	srcRect := rl.NewRectangle(
		p.frameSize*float32(int(p.frameIndex)%p.frameCount),
		p.frameSize*float32(int(p.frameIndex)%p.frameCount),
		p.frameSize,
		p.frameSize,
	)

	rl.DrawTexturePro(
		src,
		srcRect,
		p.rect,
		rl.Vector2{},
		0,
		rl.White,
	)
	rl.DrawRectangleRec(p.hitbox, rl.Black)
	return nil
}

func (p *PlayerData) animate() {
	dt := rl.GetFrameTime()
	p.frameIndex += p.frameSpeed * dt
}
