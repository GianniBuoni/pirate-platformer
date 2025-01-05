from settings import *
from sprites.animated import AnimatedSprite


class FloorSpike(AnimatedSprite):
    def __init__(self, pos, frames: list[pygame.Surface], inverted, *groups) -> None:
        if inverted:
            frames = [pygame.transform.flip(x, False, True) for x in frames]
        super().__init__(pos, frames, *groups)
        self.hitbox = self.rect.inflate(0, -32)
