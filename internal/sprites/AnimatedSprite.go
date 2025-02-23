package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	ID
	Pos
	Animation
}

func NewAnimatedSprite(obj Object, a *assets.Assets) (Sprite, error) {
	id, err := newId(obj.Image, assets.ImageLib, a)
	if err != nil {
		return nil, err
	}
	as := AnimatedSprite{
		ID:        id,
		Pos:       newPos(obj, a),
		Animation: newAnimation(),
	}
	return &as, nil
}

func (as *AnimatedSprite) Update() {}
func (as *AnimatedSprite) Draw() error {
	src, err := as.assets.GetImage(as.assetLib, as.image)
	if err != nil {
		return err
	}
	as.animate(as.rect, src)

	srcRect := rl.NewRectangle(
		as.rect.Width*float32(int(as.frameIndex)%as.frameCount),
		0,
		as.rect.Width*as.flipH,
		as.rect.Height*as.flipV,
	)

	rl.DrawTexturePro(
		src, srcRect, rl.Rectangle(*as.rect), rl.Vector2{}, 0, rl.White,
	)
	drawRect(as.hitbox, rl.Blue)
	return nil
}
