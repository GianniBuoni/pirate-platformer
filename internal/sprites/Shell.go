package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
)

type Shell struct {
	Pos
	ID
	Animation
	Player      Sprite
	attackRange float32
	canAttack   bool
	Attack      bool
}

func NewShell(obj Object, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	s := Shell{
		Pos:         newPos(obj, a),
		ID:          id,
		Animation:   newAnimation(),
		attackRange: obj.Properties.DirX,
		canAttack:   true,
	}
	return &s, nil
}

func (s *Shell) Update() {
	if s.playerInFront() {
		s.fire()
	}
	s.stop()
}

func (s *Shell) playerInFront() bool {
	pDist := s.Player.HitBox().Center().X - s.HitBox().Center().X
	pY := s.Player.HitBox().Center().Y
	pAlignedY := pY < s.HitBox().Bottom() && pY > s.HitBox().Top()

	switch s.FlipH {
	case 1:
		return pDist < s.attackRange && pDist > s.FlipH && pAlignedY
	default:
		return pDist > s.attackRange*s.FlipH && pDist < s.FlipH && pAlignedY
	}
}

func (s *Shell) SpawnFrame() bool {
	return s.Image == "shell_fire" &&
		int(s.frameIndex)%s.frameCount == 4 &&
		s.Attack == true
}

func (s *Shell) fire() {
	if s.canAttack {
		s.Image = "shell_fire"
		s.frameIndex = 0
		s.canAttack = false
		s.Attack = true
	}
}

func (s *Shell) stop() {
	if int(s.frameIndex) >= s.frameCount {
		s.Image = "shell"
		s.canAttack = true
	}
}
