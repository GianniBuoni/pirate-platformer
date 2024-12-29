from settings import *
from sprites.sprites import Sprite
from lib.timer import Timer

class Pearl(Sprite):
    def __init__(
        self, pos, collision_sprites, *groups, surf,
        bullet_direction,
        particle_funct,
        z = Z_LAYERS["main"],
    ) -> None:
        super().__init__(pos, *groups, surf=surf, z=z)
        self.collsion_sprites = [x.rect for x in collision_sprites]

        # timers
        self.timers = {
            "despawn": Timer(10000),
            "reverse": Timer(250)
        }

        # movement
        self.direction = bullet_direction
        self.speed = 100
        self.timers["despawn"].activate()

        # spawn a particle sprite
        self.particle = particle_funct

        # offset position to avoid colliison with shell
        self.rect.center = pos + vector(10 * self.direction, 7)

    def reverse(self):
        if not self.timers["reverse"]:
            self.direction *= -1
            self.timers["reverse"].activate()

    def collision(self):
        if self.rect.collidelist(self.collsion_sprites) > 0: self.destroy()

    def destroy(self):
        self.kill()
        self.particle(self.rect.center)

    def update(self, dt):
        for key in self.timers.keys(): self.timers[key].update()
        if self.timers["despawn"]:
            self.rect.centerx += self.direction * self.speed * dt
            self.collision()
        else: self.destroy()
