package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTooth(rect *Rect, path *Rect, a Assets) (*AnimatedSprite, error) {
	s, err := NewAnimatedSprite(
		"tooth", rl.NewVector2(rect.X, rect.Y), a,
		WithSpeed(PlayerSpeed*.8),
		WithDirection(rl.NewVector2(1, 0)),
	)
	s.SetHitbox(rl.Vector2{Y: (TileSize - 40) / 2}, 40, 40)
	if err != nil {
		return nil, err
	}
	s.SetPath(path)
	return s, nil
}
