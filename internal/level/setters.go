package level

import (
	"math"

	. "github.com/GianniBuoni/pirate-platformer/internal/lib"
	. "github.com/GianniBuoni/pirate-platformer/internal/sprites"
)

func (l *LevelData) AddPlayer(s Sprite) {
	p, ok := s.(*Player)
	if ok {
		// calc map limits
		top := float32(l.MapProps.TopLimit) * TileSize
		width := float32(l.Width) * TileSize
		rows := int(math.Abs(float64(l.MapProps.TopLimit))) + l.Height
		height := float32(rows) * TileSize
		cam := NewPlayerCam(p, top, width, height)
		l.player = p
		l.camera = cam
		l.AddSpriteGroup(p, l.sprites, "all")
	}
}

func (l *LevelData) AddSpriteGroup(
	s Sprite, spriteMap map[int]Sprite, groups ...string,
) {
	spriteMap[s.GetID().GID] = s
	for _, group := range groups {
		l.groups[group] = append(l.groups[group], s.GetID().GID)
	}
}
