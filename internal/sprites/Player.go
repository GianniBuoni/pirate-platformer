package sprites

import (
	"sync"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos
	ID
	Movement
	Animation
	platform Sprite
	mu       sync.Mutex
	Groups   map[string][]Sprite
	actions  map[PlayerState]bool
	cRects   map[CollisionSide]*Rect
	cSide    CollisionSide
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
		actions:   defaultStates,
		cRects:    map[CollisionSide]*Rect{},
	}
	p.getCRects()
	return &p, nil
}

func (p *Player) Update() {
	dt := rl.GetFrameTime()
	p.input(p.cSide)
	p.move(dt)
	p.Pos.Update()
	p.updateCRects()
	p.cSide = p.checkCollisionSide()
	p.image = string(p.getState(p.cSide))
}

func (p *Player) Draw() error {
	src, err := p.assets.GetImage(p.assetLib, p.image)
	if err != nil {
		return err
	}
	if p.direction.X < 0 {
		p.flipH = -1
	}
	if p.direction.X > 0 {
		p.flipH = 1
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
	for _, rect := range p.cRects {
		drawRect(rect, rl.Blue)
	}
	//drawRect(p.hitbox, rl.Green)
	return nil
}
