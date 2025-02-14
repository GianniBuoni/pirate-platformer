package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) Update() {
	p.input()
	p.move()
	p.image = string(p.getState())
	p.animate()
	p.rect.Set(
		rects.Top(p.hitbox.Top()-TileSize),
		rects.Left(p.hitbox.Left()-p.hitboxOffset),
	)
	p.oldRect.Set(
		rects.Top(p.hitbox.Top()),
		rects.Left(p.hitbox.Left()),
	)
}

func (p *PlayerData) Draw(a Assets) error {
	src, err := a.GetImage(PlayerLib, p.image)
	if err != nil {
		return err
	}

	if p.direction.X > 0 {
		p.flip = 1
	}
	if p.direction.X < 0 {
		p.flip = -1
	}

	srcRect := rl.NewRectangle(
		p.frameSize*float32(int(p.frameIndex)%p.frameCount),
		p.frameSize*float32(int(p.frameIndex)%p.frameCount),
		p.frameSize*p.flip,
		p.frameSize,
	)

	rl.DrawTexturePro(
		src,
		srcRect,
		p.rect.Rect(),
		rl.Vector2{},
		0,
		rl.White,
	)
	return nil
}

func (p *PlayerData) animate() {
	dt := rl.GetFrameTime()
	p.frameIndex += p.frameSpeed * dt
}

// call this func to debug any hitbox and rect issues
func (p *PlayerData) drawRects() {
	rl.DrawRectangleRec(p.rect.Rect(), rl.ColorAlpha(rl.Black, .25))
	rl.DrawRectangleRec(p.hitbox.Rect(), rl.ColorAlpha(rl.Green, .5))
}
