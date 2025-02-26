package sprites

import (
	"sync"
	"time"
)

type PlayerState interface {
	CheckState(PlayerAction) bool
	Timeout(PlayerAction, time.Duration)
	ToggleState(PlayerAction, bool)
}

type StateData struct {
	mu      sync.RWMutex
	actions map[PlayerAction]bool
}

func newStateData() PlayerState {
	return &StateData{
		actions: map[PlayerAction]bool{
			canAttack:   true,
			canPlatform: true,
			run:         true,
		},
	}
}

func (s *StateData) CheckState(state PlayerAction) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.actions[state]
}

func (s *StateData) ToggleState(state PlayerAction, value bool) {
	s.mu.Lock()
	s.actions[state] = value
	s.mu.Unlock()
}

func (s *StateData) Timeout(state PlayerAction, ms time.Duration) {
	s.ToggleState(state, false)
	go func() {
		timer := time.NewTimer(ms * time.Millisecond)
		<-timer.C
		s.ToggleState(state, true)
	}()
}
