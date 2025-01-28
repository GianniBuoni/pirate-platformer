package player

import (
	"github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

type Player struct {
	Attacking    bool
	Jumping      bool
	jumpDistance int
	gravity      int
	health       int
	facing       playerFacing
	state        string
	sprites.SpriteData
}

type playerFacing uint

const (
	playerDown playerFacing = iota
	playerUp
	playerLeft
	playerRight
)

type collidesWith uint

const (
	floor = iota
	left
	right
	air
)

type playerState uint

const (
	airAttack = iota
	attack
	fall
	hit
	idle
	jump
	run
	wall
)
