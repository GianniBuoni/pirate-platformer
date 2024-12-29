from settings import *
from random import choice
from lib.sprites import AnimatedSprite

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

        # movement
        self.direction = choice((-1, 1))
        self.move_speed = 200

        # collision
        self.collision_sprites: list[pygame.FRect] = [x.rect for x in collision_sprites]

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
        super().update(dt)
        self.constrain()
        self.rect.centerx += self.direction * self.move_speed * dt
