package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	frameCount int
	frameIndex float32
	FrameSpeed float32
}

func newAnimation() Animation {
	return Animation{
		FrameSpeed: FrameSpeed,
	}
}

func (a *Animation) animate(rect *Rect, src rl.Texture2D) {
	a.frameCount = int(float32(src.Width) / rect.Width)
	dt := rl.GetFrameTime()
	a.frameIndex += a.FrameSpeed * dt
}

// Call after A.animate() to prevent animation from looping.
// Passed in function is also responsible for
// toggling the current animation state to an new one.
func (a *Animation) animateOnce(
	image string, toggleFunc func(PlayerAction, bool), states ...PlayerAction,
) {
	for _, state := range states {
		if image == string(state) && int(a.frameIndex) >= a.frameCount-1 {
			a.frameIndex = 0
			if state == airAttack {
				state = attack
			}
			toggleFunc(state, false)
		}
	}
}

func (a *Animation) Draw(src rl.Texture2D, pos *Pos) {
	srcRect := rl.NewRectangle(
		pos.rect.Width*float32(int(a.frameIndex)%a.frameCount),
		0,
		pos.rect.Width*pos.FlipH,
		pos.rect.Height*pos.FlipV,
	)
	rl.DrawTexturePro(
		src, srcRect, rl.Rectangle(*pos.rect),
		rl.Vector2{}, 0, rl.White,
	)
}
