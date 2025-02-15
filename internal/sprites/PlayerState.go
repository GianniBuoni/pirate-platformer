package sprites

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerState string

const (
	airAttack PlayerState = "air_attack"
	attack    PlayerState = "attack"
	canAttack PlayerState = "can_attack"
	fall      PlayerState = "fall"
	hit       PlayerState = "hit"
	idle      PlayerState = "idle"
	jump      PlayerState = "jump"
	run       PlayerState = "run"
	wall      PlayerState = "wall"
)

func (p *PlayerData) getState() PlayerState {
	switch p.checkCollisonSide() {
	case floor:
		if p.actions[attack] {
			return attack
		}
		if p.direction == (rl.Vector2{}) {
			return idle
		} else {
			return run
		}
	case left, right:
		return wall
	case air:
		if p.actions[attack] {
			return airAttack
		}
		if p.actions[jump] {
			return jump
		}
		return fall
	default:
		return idle
	}
}
