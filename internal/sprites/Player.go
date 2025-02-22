package sprites

import (
	"github.com/GianniBuoni/pirate-platformer/internal/assets"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	ID
	Pos
	Animation
}

func NewPlayer(obj Object, a *assets.Assets) (Sprite, error) {
	id, err := newId(obj.Image, assets.PlayerLib, a)
	if err != nil {
		return nil, err
	}
	p := Player{
		ID:  id,
		Pos: newPos(obj.X, obj.Y, obj.Width, obj.Height),
	}
	return &p, nil
}

func (p *Player) Update() {}
func (p *Player) Draw() error {
	src, err := p.assets.GetImage(p.assetLib, p.image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, float32(src.Width), float32(src.Height)),
		rl.Rectangle(*p.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}
