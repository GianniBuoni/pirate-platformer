from settings import *
from random import choice

from lib.timer import Timer
from sprites.animated import AnimatedSprite

class Tooth(AnimatedSprite):
    def __init__(
        self, pos,
        frames: list[pygame.Surface],
        collision_sprites,
        *groups,
        animation_speed: float = ANIMATION_SPEED,
        z=Z_LAYERS["main"]
    ) -> None:
        super().__init__(pos, frames, *groups, animation_speed=animation_speed, z=z)
        self.reverse_timeout = Timer(250)

        # movement
        self.direction = choice((-1, 1))
        self.move_speed = 200

        # collision
        self.collision_sprites: list[pygame.FRect] = [x.rect for x in collision_sprites]

    def reverse(self):
        if not self.reverse_timeout:
            self.direction *= -1
            self.reverse_timeout.activate()

    def animate(self, dt):
        super().animate(dt)
        self.image = (
            pygame.transform.flip(self.image, True, False) if self.direction == -1
            else self.image
        )

    def constrain(self):
        l_floor_rect = pygame.FRect(self.rect.bottomleft, (-1, 1))
        r_floor_rect = pygame.FRect(self.rect.bottomright, (-1, 1))
        l_mid_rect = pygame.FRect(self.rect.midleft, (-1, 1))
        r_mid_rect = pygame.FRect(self.rect.midright, (-1, 1))

        if (
            l_floor_rect.collidelist(self.collision_sprites) == -1
            or r_floor_rect.collidelist(self.collision_sprites) == -1
            or l_mid_rect.collidelist(self.collision_sprites) > 0
            or r_mid_rect.collidelist(self.collision_sprites) > 0
        ): self.direction *= -1

    def update(self, dt):
        self.reverse_timeout.update()
        super().update(dt)
        self.constrain()
        self.rect.centerx += self.direction * self.move_speed * dt
