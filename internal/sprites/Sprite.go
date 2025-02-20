package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	image     string
	imgRect   rl.Rectangle
	rect      *Rect
	oldRect   *Rect
	hitbox    *Rect
	assetLib  AssetLibrary
	direction rl.Vector2
	speed     float32
	flipH     float32
	flipV     float32
}

var spriteDefaults BasicSprite = BasicSprite{
	imgRect:  rl.NewRectangle(0, 0, TileSize, TileSize),
	flipH:    1,
	flipV:    1,
	assetLib: ImageLib,
}

func NewSprite(
	img string, pos rl.Vector2, a Assets, opts ...func(*BasicSprite),
) (*BasicSprite, error) {
	// copy default values
	s := spriteDefaults
	// override default values with passed options
	for _, f := range opts {
		f(&s)
	}
	// check if image string is valid
	if _, err := a.GetImage(s.assetLib, img); err != nil {
		return nil, err
	}
	s.image = img
	// rects calculated
	s.rect = NewRectangle(
		pos.X, pos.Y,
		s.imgRect.Width, s.imgRect.Height,
	)
	// by default old rect and hitbox are pointers to the destRect
	s.hitbox = s.rect
	s.oldRect = s.rect

	return &s, nil
}
