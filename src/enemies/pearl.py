from settings import *
from lib.sprites import Sprite
from lib.timer import Timer


class Pearl(Sprite):
    def __init__(self, pos, *groups, surf, bullet_direction, z = Z_LAYERS["main"]) -> None:
        super().__init__(pos, *groups, surf=surf, z=z)

        self.direction = bullet_direction
        self.speed = 100
        self.despawn = Timer(10000)
        self.despawn.activate()

    def update(self, dt):
        self.despawn.update()
        if self.despawn:
            self.rect.centerx += self.direction * self.speed * dt
        else: self.kill()
