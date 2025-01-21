from random import randint

from settings import *
from sprites.animated import AnimatedSprite

class Heart(AnimatedSprite):
    def __init__(self, pos, frames: list[pygame.Surface], *groups) -> None:
        super().__init__(pos, frames, *groups)
        self.active = False

    def animate(self, dt):
        self.frames_idx += ANIMATION_SPEED * dt
        if self.frames_idx < len(self.frames):
            self.image = self.frames[int(self.frames_idx)]
        else:
            self.frames_idx = 0
            self.active = False

    def update(self, dt):
        if self.active:
            self.animate(dt)
        else:
            if randint(0, 500) == 1:
                self.active = True
