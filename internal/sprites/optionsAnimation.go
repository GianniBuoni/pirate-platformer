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
func (a *Animation) animateOnce(f func()) {
	if int(a.frameIndex) > a.frameCount {
		a.frameIndex = 0
		f()
	}
}
