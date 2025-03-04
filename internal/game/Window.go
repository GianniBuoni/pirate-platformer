package game

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowData struct {
	renderTexture rl.RenderTexture2D
	camera        rl.Camera2D
	gameScreen    rl.Rectangle
	scale         float32
}

func NewWindow() *WindowData {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.InitWindow(int32(WindowW*2), int32(WindowH*2), Title)
	rl.SetWindowMinSize(int(WindowW), int(WindowH))

	target := rl.LoadRenderTexture(int32(WindowW), int32(WindowH))
	rl.SetTextureFilter(target.Texture, rl.FilterBilinear)

	w := &WindowData{
		gameScreen:    rl.Rectangle{Width: WindowW, Height: WindowH},
		renderTexture: target,
	}
	camTarget := rl.NewVector2(w.gameScreen.Width/2, w.gameScreen.Height/2)
	w.camera = rl.NewCamera2D(camTarget, camTarget, 0, 1)
	return w
}
