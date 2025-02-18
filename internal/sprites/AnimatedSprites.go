package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedOpts interface {
	*BasicSprite | *AnimatedSprite
}

type AnimatedSprite struct {
	BasicSprite
	frameIndex float32
	frameSpeed float32
	frameCount int
}

func NewAnimatedSprite(
	img string, pos rl.Vector2, a Assets, opts ...func(*BasicSprite),
) (*AnimatedSprite, error) {
	// copy default values
	as := AnimatedSprite{
		BasicSprite: spriteDefaults,
		frameSpeed:  FrameSpeed,
	}
	// override defaults with passed options
	for _, f := range opts {
		f(&as.BasicSprite)
	}
	// check if image string is valid
	src, err := a.GetImage(as.assetLib, img)
	if err != nil {
		return nil, err
	}
	as.image = img
	// calc rects
	as.rect = NewRectangle(
		pos.X, pos.Y,
		as.imgRect.Width, as.imgRect.Height,
	)
	if as.hitbox == nil {
		as.hitbox = as.rect
		as.oldRect = as.rect
	} else {
		as.oldRect = &Rect{}
		as.oldRect.Copy(as.hitbox)
	}
	// calc frame count
	as.frameCount = int(float32(src.Width) / as.imgRect.Width)
	return &as, nil
}

func (as *AnimatedSprite) SetFrameSpeed(fs float32) {
	as.frameSpeed = fs
}
