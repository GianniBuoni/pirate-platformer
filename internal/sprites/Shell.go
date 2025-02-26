package sprites

import . "github.com/GianniBuoni/pirate-platformer/internal/lib"

type Shell struct {
	Pos
	ID
	Animation
	Player      Sprite
	attackRange float32
	canAttack   bool
	attack      bool
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
	pDist := s.Player.HitBox().Center().X - s.HitBox().Center().X
	switch s.flipH {
	case 1:
		if pDist < s.attackRange && pDist > s.flipH {
			s.fire()
			return
		}
	case -1:
		if pDist > s.attackRange && pDist < s.flipH {
			s.fire()
			return
		}
	}
	s.stop()
}

func (s *Shell) Draw() error {
	err := s.Animation.draw(s.ID, s.Pos)
	if err != nil {
		return err
	}
	return nil
}

func (s *Shell) fire() {
	if s.canAttack {
		s.attack = true
		s.frameIndex = 0
		s.image = "shell_fire"
		s.canAttack = false
	}
}

func (s *Shell) stop() {
	if int(s.frameIndex) >= s.frameCount {
		s.attack = false
		s.image = "shell"
		s.canAttack = true
	}
}
