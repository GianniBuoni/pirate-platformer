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
	p.image = string(p.getState())
	src, err := a.GetImage(PlayerLib, p.image)
	if err != nil {
		return err
	}
	p.animate(src)

	if p.direction.X > 0 {
		p.flipH = 1
	}
	if p.direction.X < 0 {
		p.flipH = -1
	}

	srcRect := rl.NewRectangle(
		p.frameSize*float32(int(p.frameIndex)%p.frameCount),
		0,
		p.frameSize*p.flipH,
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
	//p.drawRects()
	return nil
}
