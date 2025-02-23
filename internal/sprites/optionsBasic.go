package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

// The most basic sprite option.
// This defines basic sprite positions
type Pos struct {
	rect    *Rect
	oldRect *Rect
	hitbox  *Rect
	flipH   float32
	flipV   float32
}

func newPos(obj Object, a *assets.Assets) Pos {
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
		p.hitbox = NewRectangle(
			p.rect.X+hitbox.X, p.rect.Y+hitbox.Y,
			hitbox.Width, hitbox.Height,
		)
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

// Sprite info that identifies asset metadata
type ID struct {
	image    string
	assets   *assets.Assets
	assetLib assets.AssetLibrary
}

func newId(
	image string, aLib assets.AssetLibrary, a *assets.Assets,
) (ID, error) {
	if _, err := a.GetImage(aLib, image); err != nil {
		return ID{}, err
	}
	return ID{
		image:    image,
		assets:   a,
		assetLib: aLib,
	}, nil
}
