package level

import "github.com/GianniBuoni/pirate-platformer/internal/sprites"

func (l *Level) addPlayer(p *sprites.Player) {
	// assign player groups and sprites
	p.Sprites = l.spirtes
	p.Groups = l.groups

	// assign level fields
	l.camera = sprites.NewPlayerCam(p, l.Top, l.Width, l.Height)
	l.player = p
	l.AddSpriteGroup(p, l.spirtes, "all")
}
