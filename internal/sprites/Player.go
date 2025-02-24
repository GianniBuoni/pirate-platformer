package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	ID
	Pos
	Movement
	Animation
	Groups map[string][]Sprite
}

func NewPlayer(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj.Image, PlayerLib, a)
	if err != nil {
		return nil, err
	}
	p := Player{
		ID:        id,
		Pos:       newPos(obj, a),
		Movement:  newMovement(obj),
		Animation: newAnimation(),
	}
	p.SetGravity(true, 1)
	return &p, nil
}

func (p *Player) Update() {
	dt := rl.GetFrameTime()
	p.input()
	p.move(dt)
	p.Pos.Update()
}

func (p *Player) Draw() error {
	p.image = "idle"
	src, err := p.assets.GetImage(p.assetLib, p.image)
	if err != nil {
		return err
	}
	p.animate(p.rect, src)

	srcRect := rl.NewRectangle(
		p.rect.Width*float32(int(p.frameIndex)%p.frameCount),
		0,
		p.rect.Width*p.flipH,
		p.rect.Height,
	)
	rl.DrawTexturePro(
		src, srcRect,
		rl.Rectangle(*p.rect),
		rl.Vector2{}, 0, rl.White,
	)
	drawRect(p.hitbox, rl.Green)
	return nil
}
