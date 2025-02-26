package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	frameCount int
	frameIndex float32
	frameSpeed float32
}

func newAnimation() Animation {
	return Animation{
		frameSpeed: FrameSpeed,
	}
}

func (a *Animation) animate(rect *Rect, src rl.Texture2D) {
	a.frameCount = int(float32(src.Width) / rect.Width)
	dt := rl.GetFrameTime()
	a.frameIndex += a.frameSpeed * dt
}

// Call after A.animate() or A.animateEx()
// to prevent animation from looping.
// Passed in function is also responsible for
// toggling the current animation state to an new one.
func (a *Animation) animateOnce(
	image string, toggleFunc func(string), states ...string,
) {
	for _, state := range states {
		if image == state && int(a.frameIndex) >= a.frameCount-1 {
			a.frameIndex = 0
			toggleFunc(state)
		}
	}
}

func (a *Animation) draw(id ID, pos Pos) error {
	src, err := id.assets.GetImage(id.assetLib, id.image)
	if err != nil {
		return err
	}
	a.animate(pos.rect, src)

	srcRect := rl.NewRectangle(
		pos.rect.Width*float32(int(a.frameIndex)%a.frameCount),
		0,
		pos.rect.Width*pos.flipH,
		pos.rect.Height*pos.flipV,
	)
	rl.DrawTexturePro(
		src, srcRect, rl.Rectangle(*pos.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}
