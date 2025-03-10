package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraRect struct {
	level     Rect
	CamTarget rl.Vector2
	rect      *Rect
	hitbox    *Rect
	player    *Player
}

func NewPlayerCam(p *Player, top, width, height float32) *CameraRect {
	cam := CameraRect{
		level:  *NewRectangle(0, top, width, height),
		rect:   NewRectangle(0, 0, WindowW, WindowH),
		hitbox: NewRectangle(0, 0, WindowW/8, WindowH/8),
		player: p,
	}
	cam.hitbox.Set(Center(
		p.hitbox.Center().X, p.hitbox.Center().Y,
	))
	return &cam
}

func (cr *CameraRect) Update() {
	cr.playerFollow()
	cr.rect.Set(Center(
		cr.hitbox.Center().X, cr.hitbox.Center().Y,
	))
	cr.constrain()
	cr.CamTarget = rl.NewVector2(cr.rect.Center().X, cr.rect.Center().Y)
}

func (cr *CameraRect) Draw() {
	drawRect(cr.rect, rl.Red)
	drawRect(cr.hitbox, rl.Blue)
}

func (cr *CameraRect) playerFollow() {
	if cr.hitbox.Left() >= cr.player.hitbox.Left() {
		cr.hitbox.Set(Left(cr.player.hitbox.Left()))
	}
	if cr.hitbox.Right() <= cr.player.hitbox.Right() {
		cr.hitbox.Set(Right(cr.player.hitbox.Right()))
	}
	if cr.hitbox.Top() >= cr.player.hitbox.Top() {
		cr.hitbox.Set(Top(cr.player.hitbox.Top()))
	}
	if cr.hitbox.Bottom() <= cr.player.hitbox.Bottom() {
		cr.hitbox.Set(Bottom(cr.player.hitbox.Bottom()))
	}
}

func (cr *CameraRect) constrain() {
	if cr.rect.Left() <= cr.level.Left() {
		cr.rect.Set(Left(cr.level.Left()))
	}
	if cr.rect.Right() >= cr.level.Right() {
		cr.rect.Set(Right(cr.level.Right()))
	}
	if cr.rect.Top() <= cr.level.Top() {
		cr.rect.Set(Top(cr.level.Top()))
	}
	if cr.rect.Bottom() >= cr.level.Bottom() {
		cr.rect.Set(Bottom(cr.level.Bottom()))
	}
}
