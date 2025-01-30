package sprites

func (s *BasicSprite) OffsetCentre() {
	s.rect.X = s.rect.X - (s.rect.Width / 2)
	s.rect.Y = s.rect.Y - (s.rect.Height / 2)
}
