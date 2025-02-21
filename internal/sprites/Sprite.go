package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type BasicSprite struct {
	image    string
	rect     *Rect
	oldRect  *Rect
	hitbox   *Rect
	assets   *assets.Assets
	assetLib assets.AssetLibrary
	flipH    float32
	flipV    float32
}

var spriteDefaults BasicSprite = BasicSprite{
	flipH:    1,
	flipV:    1,
	assetLib: assets.ImageLib,
}

func NewSprite(obj Object, a *assets.Assets) (*BasicSprite, error) {
	// copy default values
	s := spriteDefaults

	// check if image string is valid
	if _, err := a.GetImage(s.assetLib, obj.Image); err != nil {
		return nil, err
	}
	s.image = obj.Image

	// rects calculated
	s.rect = NewRectangle(obj.X, obj.Y, obj.Width, obj.Height)

	// by default old rect and hitbox are pointers to the destRect
	s.hitbox = s.rect
	s.oldRect = s.rect

	return &s, nil
}
