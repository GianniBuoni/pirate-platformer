package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *PlayerData) animate(src rl.Texture2D) {
	dt := rl.GetFrameTime()
	p.frameCount = int(float32(src.Width) / p.imgRect.Width)
	p.frameIndex += p.frameSpeed * dt
	switch p.image {
	case string(jump):
		p.animateOnce(jump)
	case string(attack), string(airAttack):
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
