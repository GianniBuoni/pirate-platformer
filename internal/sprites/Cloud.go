package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cloud struct {
	Pos
	ID
	Movement
	spawnX float32
}

func NewCloud(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	id, err := newId(obj, ImageLib, a)
	if err != nil {
		return nil, err
	}
	c := &Cloud{
		ID:       id,
		Pos:      newPos(obj, a),
		Movement: newMovement(obj),
		spawnX:   obj.Properties.Lifetime * TileSize,
	}
	c.rect.Set(Bottom(obj.Y))
	return c, nil
}

func (c *Cloud) Update() {
	c.MoveX(c.rect, rl.GetFrameTime())
	if c.rect.Right() <= 0 {
		c.rect.Set(Left(c.spawnX))
	}
}

func (c *Cloud) Draw(*ID, *Pos) error {
	src, err := c.assets.GetImage(c.assetLib, c.Image)
	if err != nil {
		return err
	}
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, c.rect.Width, c.rect.Height),
		rl.Rectangle(*c.rect),
		rl.Vector2{}, 0, rl.White,
	)
	return nil
}
