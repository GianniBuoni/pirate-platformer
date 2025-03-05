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
	FlipH        float32
	FlipV        float32
}

func newPos(obj Object, a *Assets) Pos {
	p := Pos{
		rect:    NewRectangle(obj.X, obj.Y, obj.Width, obj.Height),
		oldRect: NewRectangle(obj.X, obj.Y, obj.Width, obj.Height),
	}
	if obj.Properties.FlipH == 0 {
		p.FlipH = 1
	} else {
		p.FlipH = obj.Properties.FlipH
	}
	if obj.Properties.FlipV == 0 {
		p.FlipV = 1
	} else {
		p.FlipV = obj.Properties.FlipV
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

func (p *Pos) GetPos() *Pos {
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

func (p *Pos) Facing() float32 {
	return p.FlipH
}

func (p *Pos) Update() error {
	p.rect.X = p.hitbox.X - p.hitboxOffset.X
	p.rect.Y = p.hitbox.Y - p.hitboxOffset.Y
	return nil
}

// Sprite info that identifies asset metadata
type ID struct {
	Src      rl.Texture2D
	Image    string
	GID      int
	assets   *Assets
	assetLib AssetLibrary
	Kill     bool
}

func newId(
	obj Object, aLib AssetLibrary, a *Assets,
) (ID, error) {
	id := ID{
		Image:    obj.Image,
		GID:      obj.TiledID,
		assets:   a,
		assetLib: aLib,
	}
	var err error
	id.Src, err = a.GetImage(aLib, obj.Image)
	if err != nil {
		return ID{}, err
	}
	return id, nil
}

func (id *ID) GetID() *ID {
	return id
}
