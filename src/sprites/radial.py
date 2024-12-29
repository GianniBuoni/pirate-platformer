from settings import *
from .sprites import Sprite

from math import sin, cos, radians

class RadialSprite(Sprite):
    def __init__(
        self, *groups, pos, radius, speed, start_angle, end_angle,
        surf, z= Z_LAYERS["main"]
    ) -> None:
        self.center = pos
        self.radius = radius
        self.start_angle = start_angle
        self.end_angle = end_angle
        self.angle = self.start_angle

        # movement
        self.speed = speed
        self.direction = 1
        self.constrained = False if self.end_angle == -1 else True

        super().__init__(self.get_pos(), *groups, surf=surf, z=z)

    def get_pos(self) -> tuple:
        x = self.center[0] + cos(radians(self.angle)) * self.radius # adjacent
        y = self.center[1] + sin(radians(self.angle)) * self.radius # opponent
        return(x,y)

    def constrain(self):
        if self.constrained:
            if self.angle >= self.end_angle: self.direction = -1
            if self.angle <= self.start_angle: self.direction = 1

    def update(self, dt) -> None:
        self.angle += self.direction * self.speed * dt
        self.constrain()
        self.rect.center = self.get_pos()
