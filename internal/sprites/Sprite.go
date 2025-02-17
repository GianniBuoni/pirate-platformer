package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	image        string
	imgRect      rl.Rectangle
	rect         SpriteRect
	oldRect      SpriteRect
	hitbox       SpriteRect
	assetLib     AssetLibrary
	direction    rl.Vector2
	hitboxOffset float32
	speed        float32
	flipH        float32
	flipV        float32
}

func NewSprite(
	img string, pos rl.Vector2, a Assets, opts ...func(*BasicSprite),
) (*BasicSprite, error) {
	// default values
	s := &BasicSprite{
		image:    img,
		imgRect:  rl.NewRectangle(0, 0, TileSize, TileSize),
		flipH:    1,
		flipV:    1,
		assetLib: ImageLib,
	}
	// override default values with passed options
	for _, f := range opts {
		f(s)
	}
	// check if image string is valid
	if _, err := a.GetImage(s.assetLib, img); err != nil {
		return &BasicSprite{}, err
	}
	// rects calculated
	s.rect = rects.NewRectangle(
		pos.X, pos.Y,
		s.imgRect.Width, s.imgRect.Height,
	)
	// by default old rect and hitbox are pointers to the destRect
	s.hitbox = s.rect
	s.oldRect = s.rect

	return s, nil
}

func WithImgPos(pos rl.Vector2) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.X = pos.X
		bs.imgRect.Y = pos.Y
	}
}

func WithImgWidth(w float32) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.Width = w
	}
}

func WithImgHeight(h float32) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.Height = h
	}
}

func WithAssetLib(al AssetLibrary) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.assetLib = al
	}
}

func WithFlipV(b bool) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		if b {
			bs.flipV = -1
		}
	}
}

func WithFlipH(b bool) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		if b {
			bs.flipH = -1
		}
	}
}
