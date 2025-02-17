package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	"github.com/GianniBuoni/pirate-platformer/internal/rects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicSprite struct {
	image     string
	imgRect   rl.Rectangle
	rect      SpriteRect
	oldRect   SpriteRect
	hitbox    SpriteRect
	direction rl.Vector2
	speed     float32
	flipH     float32
	flipV     float32
}

func NewSprite(
	img string, pos rl.Vector2, a Assets, opts ...func(*BasicSprite),
) (*BasicSprite, error) {
	// check if image string is valid
	if _, err := a.GetImage(ImageLib, img); err != nil {
		return &BasicSprite{}, err
	}
	s := &BasicSprite{
		image: img,
		flipH: 1,
		flipV: 1,
	}
	// sprite defaults to TileSize x TileSize rect
	s.imgRect = rl.NewRectangle(0, 0, TileSize, TileSize)
	// override defaults with any added options
	for _, f := range opts {
		f(s)
	}
	// dest rect calculated after sprite data processed
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
