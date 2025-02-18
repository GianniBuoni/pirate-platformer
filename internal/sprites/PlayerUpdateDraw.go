package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) Update() {
	p.input()
	p.move()
	p.rect.Set(
		Top(p.hitbox.Top()-p.hitboxOffset.Y),
		Left(p.hitbox.Left()-p.hitboxOffset.X),
	)
	p.oldRect.Copy(p.hitbox)
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
		p.imgRect.X*float32(int(p.frameIndex)%p.frameCount),
		0,
		p.imgRect.Width*p.flipH,
		p.imgRect.Height,
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
