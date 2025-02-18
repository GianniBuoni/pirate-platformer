package sprites

func (as *AnimatedSprite) Update() {
	as.oldRect.Copy(as.hitbox)
	as.Movement(as.rect)
	as.constrain()
}

func (as *AnimatedSprite) constrain() {
	switch as.direction.Y {
	// horizontal check
	case 0:
		if as.rect.X <= as.pathRect.Left() {
			as.direction.X = 1
		}
		if as.rect.X >= as.pathRect.Right() {
			as.direction.X = -1
		}
	default:
		if as.rect.Y <= as.pathRect.Top() {
			as.direction.Y = 1
		}
		if as.rect.Y >= as.pathRect.Bottom() {
			as.direction.Y = -1
		}
	}
}
