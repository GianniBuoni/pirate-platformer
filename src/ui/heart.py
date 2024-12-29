from settings import *
from sprites.animated import AnimatedSprite

class Heart(AnimatedSprite):
    def __init__(self, pos, frames: list[pygame.Surface], *groups) -> None:
        super().__init__(pos, frames, *groups)
