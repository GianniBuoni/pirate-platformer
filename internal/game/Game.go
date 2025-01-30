package game

import . "github.com/GianniBuoni/pirate-platformer/internal/interfaces"

type GameData struct {
  running bool
}

func NewGame() Game {
  return &GameData{
    running: true,
  }
}
