package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShellSprite struct {
	AnimatedSprite
	player *Rect
	attack bool
}

func NewShell(rect *Rect, a Assets, flip bool) (*ShellSprite, error) {
	s := ShellSprite{
		AnimatedSprite: AnimatedSprite{
			BasicSprite: BasicSprite{
				image:    "shell",
				assetLib: ImageLib,
				hitbox: NewRectangle(
					PlatHitbox, 16, TileSize-PlatHitbox*2, TileSize-32,
				),
				oldRect: &Rect{},
				flipH:   1,
			},
			frameSpeed: FrameSpeed,
		},
	}
	if _, err := a.GetImage(s.assetLib, s.image); err != nil {
		return nil, err
	}
	// CALC RECTS
	s.rect = rect
	s.hitbox.Set(Center(s.rect.Center().X, s.rect.Center().Y))
	s.oldRect.Copy(s.hitbox)
	if flip {
		s.flipH = -1
	}
	return &s, nil
}

func (s *ShellSprite) Update() {
	var playerInFront bool
	if s.flipH == 1 {
		playerInFront = s.rect.Right() <= s.player.Right()
	} else {
		playerInFront = s.rect.Left() >= s.player.Left()
	}
	if rl.CheckCollisionRecs(s.rect.Rect(), s.player.Rect()) && playerInFront {
		s.attack = true
	}
}

func (s *ShellSprite) IsAttacking() bool {
	if int(s.frameIndex) == 4 {
		return s.attack
	}
	return false
}

func (s *ShellSprite) Draw(a Assets) error {
	if s.attack {
		s.image = "shell_fire"
	} else {
		s.image = "shell"
	}

	src, err := a.GetImage(s.assetLib, s.image)
	if err != nil {
		return err
	}
	s.frameCount = int(float32(src.Width) / s.rect.Width)
	s.frameIndex += s.frameSpeed * rl.GetFrameTime()
	if int(s.frameIndex) >= s.frameCount-1 {
		s.frameIndex = 0
		s.attack = false
	}

	srcRect := rl.NewRectangle(
		s.rect.Width*float32(int(s.frameIndex)%s.frameCount),
		0,
		s.rect.Width*s.flipH,
		s.rect.Height,
	)
	rl.DrawTexturePro(src, srcRect, s.rect.Rect(), rl.Vector2{}, 0, rl.White)
	//s.drawRects(rl.Blue)
	return nil
}

func (s *ShellSprite) SetPlayer(p Sprite) {
	s.player = p.Rect()
}
