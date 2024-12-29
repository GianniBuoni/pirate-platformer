from settings import *
from .animated import AnimatedSprite

class Item(AnimatedSprite):
    def __init__(
        self, pos,
        frames: list[pygame.Surface],
        *groups,
        animation_speed: float = ANIMATION_SPEED,
        z = Z_LAYERS["main"]
    ) -> None:

        super().__init__(pos, frames, *groups, animation_speed=animation_speed, z=z)
        self.rect.center = pos
