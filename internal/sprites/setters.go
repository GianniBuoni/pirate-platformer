package sprites

func (s *BasicSprite) OffsetCentre() {
	s.rect.X = s.rect.X - (s.rect.Width / 2)
	s.rect.Y = s.rect.Y - (s.rect.Height / 2)
}

func (p *PlayerData) SetLeft(x float32) {
	p.hitbox.X = x
}

func (p *PlayerData) SetRight(x float32) {
	p.hitbox.X = x - p.hitbox.Width
}

func (p *PlayerData) SetTop(y float32) {
	p.hitbox.Y = y
}

func (p *PlayerData) SetBottom(y float32) {
	p.hitbox.Y = y - p.hitbox.Height
}
