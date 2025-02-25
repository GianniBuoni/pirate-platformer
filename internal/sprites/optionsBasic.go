package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// The most basic sprite option.
// This defines basic sprite positions
type Pos struct {
	rect         *Rect
	oldRect      *Rect
	hitbox       *Rect
	hitboxOffset rl.Vector2
	flipH        float32
	flipV        float32
}

func newPos(obj Object, a *Assets) Pos {
	p := Pos{
		rect:    NewRectangle(obj.X, obj.Y, obj.Width, obj.Height),
		oldRect: NewRectangle(obj.X, obj.Y, obj.Width, obj.Height),
	}
	if obj.Properties.FlipH == 0 {
		p.flipH = 1
	} else {
		p.flipH = obj.Properties.FlipH
	}
	if obj.Properties.FlipV == 0 {
		p.flipV = 1
	} else {
		p.flipV = obj.Properties.FlipV
	}
	hitbox, ok := a.Hitboxes[obj.Image]
	if !ok {
		p.hitbox = p.rect
	} else {
		p.hitboxOffset = rl.NewVector2(hitbox.X, hitbox.Y)
		p.hitbox = NewRectangle(
			p.rect.X+hitbox.X, p.rect.Y+hitbox.Y,
			hitbox.Width, hitbox.Height,
		)
		p.oldRect.Copy(p.hitbox)
	}
	return p
}

func (p *Pos) Rect() *Rect {
	return p.rect
}

func (p *Pos) HitBox() *Rect {
	return p.hitbox
}

func (p *Pos) OldRect() *Rect {
	return p.oldRect
}

func (p *Pos) Update() {
	p.rect.X = p.hitbox.X - p.hitboxOffset.X
	p.rect.Y = p.hitbox.Y - p.hitboxOffset.Y
	p.oldRect.Copy(p.hitbox)
}

// Sprite info that identifies asset metadata
type ID struct {
	image    string
	id       int
	assets   *Assets
	assetLib AssetLibrary
}

func newId(
	obj Object, aLib AssetLibrary, a *Assets,
) (ID, error) {
	if _, err := a.GetImage(aLib, obj.Image); err != nil {
		return ID{}, err
	}
	return ID{
		image:    obj.Image,
		id:       obj.Id,
		assets:   a,
		assetLib: aLib,
	}, nil
}
