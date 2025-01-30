package window

import rl "github.com/gen2brain/raylib-go/raylib"

func (w *WindowData) GetScreenRect() rl.Rectangle {
  return w.gameScreen
}

func (w *WindowData) GetRenderTexture() rl.RenderTexture2D {
  return w.renderTexture
}
