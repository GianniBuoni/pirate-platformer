package lib

type Stats struct {
	CurrentLevel  int
	UnlockedLevel int
	Coins         int
	playerHealth  int
	maxHealth     int
	Paused        bool
}

func NewStats() *Stats {
	return &Stats{
		CurrentLevel:  1,
		UnlockedLevel: 4,
		Coins:         0,
		playerHealth:  3,
		maxHealth:     3,
	}
}

func (s *Stats) PlayerHP() int {
	return s.playerHealth
}

func (s *Stats) SetMaxHP(hpTpAdd int) {
	maxHP := s.maxHealth + hpTpAdd
	s.maxHealth = min(maxHP, MaxHealth)
}

func (s *Stats) AddHP(hpToAdd int) {
	health := s.playerHealth + hpToAdd
	s.playerHealth = max(0, min(health, s.maxHealth))
}
