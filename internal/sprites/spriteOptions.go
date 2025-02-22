package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
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

func newPos(x, y, width, height float32) Pos {
	p := Pos{
		rect:    NewRectangle(x, y, width, height),
		oldRect: NewRectangle(x, y, width, height),
	}
	p.hitbox = p.rect
	return p
}

func (p *Pos) SetHitbox(width, height float32) {
	p.hitbox = NewRectangle(0, 0, width, height)
	p.hitbox.Set(Center(p.rect.Center().X, p.rect.Center().Y))
	p.oldRect.Copy(p.hitbox)
}

// call this func in the draw method to debug any hitbox and rect issues
func (p *Pos) drawHitbox(c rl.Color) {
	rl.DrawRectangleRec(
		rl.Rectangle(*p.hitbox), rl.ColorAlpha(c, .5),
	)
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
