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
	canAttack:   true,
	canPlatform: true,
	run:         true,
}

func (p *Player) getState(side CollisionSide) PlayerState {
	if p.actions[hit] {
		p.actions[attack] = false
		p.actions[jump] = false
		return hit
	}

	if p.actions[attack] {
		p.actions[jump] = false
	}

	if p.actions[jump] {
		return jump
	}

	switch side {
	case floor:
		p.actions[wall] = false
		if p.actions[attack] {
			return attack
		}
		if p.direction.X != 0 {
			return run
		}
		return idle
	case left, right:
		if p.actions[wall] {
			p.SetGravity(false, .8)
			return wall
		}
		return fall
	case air:
		p.SetGravity(true, 1)
		p.platform = nil
		if p.actions[attack] {
			return airAttack
		}
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

func (p *Player) toggleState(state string) {
	if state == string(airAttack) {
		state = string(attack)
	}
	p.actions[PlayerState(state)] = false
}
