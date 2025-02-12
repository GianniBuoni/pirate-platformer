package sprites

import (
	"fmt"

	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSprite struct {
	BasicSprite
	srcRect rl.Rectangle
}

func NewTileSprite(
	image string,
	pos rl.Vector2,
	srcPos rl.Vector2,
	a Assets,
) (Sprite, error) {
	_, err := a.GetImage(TilesetLib, image)
	if err != nil {
		return nil, fmt.Errorf(
			"New tile sprite %s, could not be created. %w", image, err,
		)
	}

	sprite := &TileSprite{}
	sprite.image = image
	sprite.rect = rl.Rectangle{
		X: pos.X, Y: pos.Y,
		Width: TileSize, Height: TileSize,
	}
	sprite.srcRect = rl.Rectangle{
		X: srcPos.X, Y: srcPos.Y,
		Width: TileSize, Height: TileSize,
	}
	return sprite, nil
}

func (s *TileSprite) Draw(a Assets) error {
	src, err := a.GetImage(TilesetLib, s.image)
	if err != nil {
		return err
	}

	rl.DrawTexturePro(
		src, s.srcRect, s.rect, rl.Vector2{}, 0, rl.White,
	)

	return nil
}
