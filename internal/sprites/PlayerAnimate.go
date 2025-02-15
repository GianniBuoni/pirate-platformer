package sprites

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) animate(src rl.Texture2D) {
	dt := rl.GetFrameTime()
	p.frameCount = int(float32(src.Width) / p.frameSize)
	p.frameIndex += p.frameSpeed * dt
	switch p.image {
	case string(jump):
		p.animateOnce(jump)
	case string(attack), string(airAttack):
		fmt.Println(p.image)
		p.animateOnce(attack)
	}
}

func (p *PlayerData) animateOnce(state PlayerState) {
	if int(p.frameIndex) >= p.frameCount-1 {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.frameIndex = 0
		p.actions[state] = false
	}
}

// call this func in the draw method to debug any hitbox and rect issues
func (p *PlayerData) drawRects() {
	rl.DrawRectangleRec(p.rect.Rect(), rl.ColorAlpha(rl.Black, .25))
	rl.DrawRectangleRec(p.hitbox.Rect(), rl.ColorAlpha(rl.Green, .5))
}
