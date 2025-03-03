package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cloud struct {
	ID
	Pos
	Movement
	spawnX float32
}

func NewCloud(obj Object, aLib AssetLibrary, a *Assets) (Sprite, error) {
	c := &Cloud{
		Pos:      newPos(obj, a),
		Movement: newMovement(obj),
		spawnX:   obj.Properties.Lifetime * TileSize,
	}
	var err error
	c.ID, err = newId(obj, aLib, a)
	if err != nil {
		return nil, err
	}
	c.rect.Set(Bottom(obj.Y))
	return c, nil
}

func (c *Cloud) Update() error {
	c.MoveX(c.rect, rl.GetFrameTime())
	if c.rect.Right() <= 0 {
		c.rect.Set(Left(c.spawnX))
	}
	return nil
}

func (c *Cloud) Draw(src rl.Texture2D, pos *Pos) {
	rl.DrawTexturePro(
		src,
		rl.NewRectangle(0, 0, c.rect.Width, c.rect.Height),
		rl.Rectangle(*c.rect),
		rl.Vector2{}, 0, rl.White,
	)
}
