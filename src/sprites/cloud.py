from .sprites import Sprite
from settings import *
from random import randint

class Cloud(Sprite):
    def __init__(self, pos, *groups, surf, z = Z_LAYERS["bg"]) -> None:
        super().__init__(pos, *groups, surf=surf, z=z)
        self.rect.bottomleft = pos
        self.direction = -1
        self.speed = randint(100, 150)

    def update(self, dt):
        self.rect.centerx += self.direction * self.speed * dt
        if self.rect.right <= 0: self.kill()
