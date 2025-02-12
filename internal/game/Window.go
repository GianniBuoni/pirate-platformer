package game

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowData struct {
	scale         float32
	gameScreen    rl.Rectangle
	renderTexture rl.RenderTexture2D
}

func NewWindow() *WindowData {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.InitWindow(int32(WindowW), int32(WindowH), Title)
	rl.SetWindowMinSize(int(WindowW), int(WindowH))

	target := rl.LoadRenderTexture(int32(WindowW), int32(WindowH))
	rl.SetTextureFilter(target.Texture, rl.FilterBilinear)

	return &WindowData{
		gameScreen:    rl.Rectangle{Width: WindowW, Height: WindowH},
		renderTexture: target,
	}
}
