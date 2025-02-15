package sprites

import "fmt"

func (p *PlayerData) attack(c CollisionSide) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.actions["attack"] = true
	switch c {
	case air:
		fmt.Println("air attack")
	default:
		fmt.Println("attack")
	}
}
