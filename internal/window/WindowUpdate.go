package window

import rl "github.com/gen2brain/raylib-go/raylib"

func (w *WindowData) Update() {
	w.getWindowScale()
}

func (w *WindowData) getWindowScale() {
	w.scale = min(float32(rl.GetScreenWidth())/float32(w.gameScreen.Width),
		float32(rl.GetScreenHeight())/float32(w.gameScreen.Height))
}
