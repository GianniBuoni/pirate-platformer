package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (as *AnimatedSprite) Draw(assets Assets) error {
	dt := rl.GetFrameTime()
	src, err := assets.GetImage(ImageLib, as.image)
	if err != nil {
		return err
	}
	// update frame info
	as.frameIndex += as.frameSpeed * dt
	srcRect := rl.NewRectangle(
		as.imgRect.Width*float32(int(as.frameIndex)%as.frameCount),
		as.imgRect.Y,
		as.imgRect.Width*as.flipH,
		as.imgRect.Height*as.flipV,
	)
	rl.DrawTexturePro(
		src, srcRect, as.rect.Rect(), rl.Vector2{}, 0, rl.White,
	)
	return nil
}
