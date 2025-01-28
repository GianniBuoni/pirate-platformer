package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite interface {
	Update()
	Draw()
}

type MovingSprite interface {
	Sprite
	Move()
	Collide()
}

type AnimatedSprite interface {
	MovingSprite
	Animate()
}

type SpriteData struct {
	moveSpeed    float32
	frameIdx     float32
	frameCount   int
	animateSpeed int
	direction    rl.Vector2
	srcGrid      rl.Vector2
	srcImage     string
	rect         rl.Rectangle
	OldRect      rl.Rectangle
	Hitbox       rl.Rectangle
}
