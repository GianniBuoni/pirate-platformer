package main

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
  rl.InitWindow(WindowW, WindowH, Title)
  defer rl.CloseWindow()

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Pink)
    rl.EndDrawing()
  }
}
