package interfaces

import rl "github.com/gen2brain/raylib-go/raylib"

type Window interface {
	Update()
	Draw()

	// getters
	GetScreenRect() rl.Rectangle
	GetRenderTexture() rl.RenderTexture2D
}
