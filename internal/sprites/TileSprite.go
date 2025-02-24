package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSprite struct {
	ID
	Pos
	imgPos rl.Vector2
}

func NewTileSprite(tile Tile, a *Assets) (
	Sprite, error,
) {
	id, err := newId(tile.Image, TilesetLib, a)
	if err != nil {
		return nil, err
	}

	obj := Object{
		X: tile.X, Y: tile.Y, Width: TileSize, Height: TileSize,
	}

	ts := TileSprite{
		ID:     id,
		Pos:    newPos(obj, a),
		imgPos: rl.NewVector2(tile.ImgX, tile.ImgY),
	}
	return &ts, nil
}

func (ts *TileSprite) Update() {}

func (ts *TileSprite) Draw() error {
	src, err := ts.assets.GetImage(ts.assetLib, ts.image)
	if err != nil {
		return err
	}
	srcRect := rl.NewRectangle(
		ts.imgPos.X, ts.imgPos.Y, TileSize, TileSize,
	)
	rl.DrawTexturePro(
		src, srcRect, rl.Rectangle(*ts.rect),
		rl.Vector2{}, 0, rl.White,
	)
	drawRect(ts.hitbox, rl.Red)
	return nil
}
