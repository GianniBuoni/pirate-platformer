package sprites

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

func (p *Player) getState() {
	var state PlayerState
	if p.direction.X == 0 {
		state = idle
	} else {
		state = run
		return
	}
	p.image = string(state)
}
