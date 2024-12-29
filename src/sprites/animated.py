from settings import *
from .sprites import Sprite

class AnimatedSprite(Sprite):
    def __init__(
            self, pos, frames: list[pygame.Surface], *groups,
            animation_speed: float = ANIMATION_SPEED,
            z = Z_LAYERS["main"]
    ) -> None:
        self.frames, self.frames_idx = frames, 0
        super().__init__(pos, *groups, surf = self.frames[self.frames_idx], z=z)
        self.animation_speed = animation_speed

    def animate(self, dt):
        self.frames_idx += self.animation_speed * dt
        self.image = self.frames[int(self.frames_idx % len(self.frames))]

    def update(self, dt):
        self.animate(dt)
