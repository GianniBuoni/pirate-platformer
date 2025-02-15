package sprites

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) animate() {
	p.image = string(p.getState())

	dt := rl.GetFrameTime()
	p.frameIndex += p.frameSpeed * dt
	switch p.image {
	case "jump":
		if int(p.frameIndex) >= p.frameCount {
			p.mu.Lock()
			defer p.mu.Unlock()
			p.frameIndex = 0
			p.actions["jump"] = false
			p.image = "fall"
		}
		fmt.Println(p.frameIndex, p.frameCount)
	}
}

// call this func in the draw method to debug any hitbox and rect issues
func (p *PlayerData) drawRects() {
	rl.DrawRectangleRec(p.rect.Rect(), rl.ColorAlpha(rl.Black, .25))
	rl.DrawRectangleRec(p.hitbox.Rect(), rl.ColorAlpha(rl.Green, .5))
}
