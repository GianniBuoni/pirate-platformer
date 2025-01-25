package lib

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Assets struct {
	mu     sync.RWMutex
	Frames map[string][]rl.Texture2D
}
