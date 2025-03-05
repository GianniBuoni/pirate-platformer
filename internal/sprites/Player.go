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
	Groups   map[string][]int
	Sprites  map[int]Sprite
	cRects   map[CollisionSide]*Rect
	cSide    CollisionSide
}

func NewPlayer(obj Object, stats *Stats, a *Assets) (Sprite, error) {
	p := Player{
		Pos:       newPos(obj, a),
		Movement:  newMovement(obj),
		Animation: newAnimation(),
		state:     newStateData(),
		stats:     stats,
		cRects:    map[CollisionSide]*Rect{},
	}
	var err error
	p.ID, err = newId(obj, PlayerLib, a)
	if err != nil {
		return nil, err
	}
	p.getCRects()
	return &p, nil
}

func (p *Player) Update() error {
	dt := rl.GetFrameTime()
	p.oldRect.Copy(p.hitbox)
	p.input(p.cSide)
	p.move(dt)
	p.Pos.Update()
	p.updateCRects()
	p.cSide = p.checkCollisionSide()
	p.Image = string(p.getState())

	var err error
	p.Src, err = p.assets.GetImage(p.assetLib, p.Image)
	if err != nil {
		return err
	}
	if p.direction.X < 0 {
		p.FlipH = -1
	}
	if p.direction.X > 0 {
		p.FlipH = 1
	}
	p.animate(p.rect, p.Src)
	p.animateOnce(
		p.Image, p.state.ToggleState, airAttack, attack, hit, jump,
	)
	return nil
}

func (p *Player) Draw(src rl.Texture2D, pos *Pos) {
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
}
