package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos
	ID
	Movement
	Animation
	platform Sprite
	state    PlayerState
	stats    *Stats
	Groups   map[string][]Sprite
	cRects   map[CollisionSide]*Rect
	cSide    CollisionSide
}

func NewPlayer(obj Object, stats *Stats, a *Assets) (Sprite, error) {
	id, err := newId(obj, PlayerLib, a)
	if err != nil {
		return nil, err
	}
	p := Player{
		ID:        id,
		Pos:       newPos(obj, a),
		Movement:  newMovement(obj),
		Animation: newAnimation(),
		state:     newStateData(),
		stats:     stats,
		cRects:    map[CollisionSide]*Rect{},
	}
	p.getCRects()
	return &p, nil
}

func (p *Player) Update() {
	dt := rl.GetFrameTime()
	p.oldRect.Copy(p.hitbox)
	p.input(p.cSide)
	p.move(dt)
	p.Pos.Update()
	p.updateCRects()
	p.cSide = p.checkCollisionSide()
	p.Image = string(p.getState())
}

func (p *Player) Draw(id *ID, pos *Pos) error {
	src, err := p.assets.GetImage(p.assetLib, p.Image)
	if err != nil {
		return err
	}
	if p.direction.X < 0 {
		p.FlipH = -1
	}
	if p.direction.X > 0 {
		p.FlipH = 1
	}
	p.animate(p.rect, src)
	p.animateOnce(
		p.Image, p.state.ToggleState, airAttack, attack, hit, jump,
	)

	srcRect := rl.NewRectangle(
		p.rect.Width*float32(int(p.frameIndex)%p.frameCount),
		0,
		p.rect.Width*p.FlipH,
		p.rect.Height,
	)
	rl.DrawTexturePro(
		src, srcRect,
		rl.Rectangle(*p.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}
