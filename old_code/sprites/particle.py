from settings import *
from sprites.animated import AnimatedSprite


class Particle(AnimatedSprite):
    def __init__(
        self, pos, frames: list[pygame.Surface],
        *groups,
        animation_speed: float = ANIMATION_SPEED,
        z = Z_LAYERS["fg"]
    ) -> None:
        super().__init__(pos, frames, *groups, animation_speed=animation_speed, z=z)
        self.rect.center = pos

    def animate(self, dt):
        self.frames_idx += self.animation_speed * dt
        if self.frames_idx < len(self.frames):
            self.image = self.frames[int(self.frames_idx)]
        else:
            self.kill()
