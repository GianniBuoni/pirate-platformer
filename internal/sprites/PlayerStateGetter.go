package sprites

type PlayerAction string

const (
	airAttack   PlayerAction = "air_attack"
	attack      PlayerAction = "attack"
	canAttack   PlayerAction = "can_attack"
	fall        PlayerAction = "fall"
	hit         PlayerAction = "hit"
	idle        PlayerAction = "idle"
	jump        PlayerAction = "jump"
	canPlatform PlayerAction = "can_platform"
	run         PlayerAction = "run"
	wall        PlayerAction = "wall"
)

func (p *Player) getState() PlayerAction {
	// check mutally exclusive input states
	if p.state.CheckState(hit) {
		p.state.ToggleState(attack, false)
		p.state.ToggleState(jump, false)
		return hit
	}
	if p.state.CheckState(attack) {
		p.state.ToggleState(jump, false)
	}
	if p.state.CheckState(jump) {
		return jump
	}

	// check collision side
	switch p.cSide {
	case floor:
		p.state.ToggleState(wall, false)
		if p.state.CheckState(attack) {
			return attack
		}
		if p.direction.X != 0 {
			return run
		}
		return idle
	case left, right:
		if p.state.CheckState(wall) {
			p.SetGravity(false, 0.8)
			return wall
		}
	case air:
		p.SetGravity(true, 1.0)
		p.platform = nil
		if p.state.CheckState(attack) {
			return airAttack
		}
		return fall
	}
	return idle
}
