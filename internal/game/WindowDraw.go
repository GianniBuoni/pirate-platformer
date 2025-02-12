package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (w *WindowData) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black) // letterbox color

	// Draw render texture to screen, properly scaled
	rl.DrawTexturePro(
		w.renderTexture.Texture,
		rl.Rectangle{Width: float32(w.renderTexture.Texture.Width), Height: float32(-w.renderTexture.Texture.Height)},
		rl.Rectangle{
			X:      (float32(rl.GetScreenWidth()) - w.gameScreen.Width*w.scale) * 0.5,
			Y:      (float32(rl.GetScreenHeight()) - w.gameScreen.Height*w.scale) * 0.5,
			Width:  w.gameScreen.Width * w.scale,
			Height: w.gameScreen.Height * w.scale,
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
	rl.EndDrawing()
}
