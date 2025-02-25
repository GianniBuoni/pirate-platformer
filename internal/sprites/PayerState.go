package sprites

import (
	"time"
)

type PlayerState string

const (
	airAttack   PlayerState = "air_attack"
	attack      PlayerState = "attack"
	canAttack   PlayerState = "can_attack"
	fall        PlayerState = "fall"
	hit         PlayerState = "hit"
	idle        PlayerState = "idle"
	jump        PlayerState = "jump"
	canPlatform PlayerState = "can_platform"
	run         PlayerState = "run"
	wall        PlayerState = "wall"
)

var defaultStates = map[PlayerState]bool{
	run: true,
}

func (p *Player) getState(side CollisionSide) PlayerState {
	switch side {
	case floor:
		p.actions[wall] = false
		if p.direction.X == 0 {
			return idle
		} else {
			return run
		}
	case left, right:
		if p.actions[wall] {
			p.SetGravity(false, .8)
			return wall
		}
		return fall
	case air:
		p.SetGravity(true, 1)
		p.platform = nil
		return fall
	default:
		return idle
	}
}

func (p *Player) timeout(state PlayerState, ms time.Duration) {
	p.actions[state] = false
	go func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		timer := time.NewTimer(ms * time.Millisecond)
		<-timer.C
		p.actions[state] = true
	}()
}
