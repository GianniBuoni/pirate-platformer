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

func NewAnimation() Animation {
	return Animation{
		frameSpeed: FrameSpeed,
	}
}

func (a *Animation) animate(dt float32) {
	a.frameIndex += a.frameSpeed * dt
}

// Call A.animateEx() if A.frameCount is variable.
// Most likely to be needed if sprite canges states.
func (a *Animation) animateEx(dt, objWidth float32, src rl.Texture2D) {
	a.frameCount = int(float32(src.Width) / objWidth)
	a.animate(dt)
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
