from settings import *
from .sprites import Sprite

from math import sin, cos, tan, radians

class RadialSprite(Sprite):
    def __init__(
        self, *groups, pos, radius, speed, start_angle, end_angle,
        surf, z= Z_LAYERS["main"]
    ) -> None:
        self.center = pos
        self.radius = radius
        self.speed = speed
        self.start_angle = start_angle
        self.end_angle = end_angle
        self.angle = self.start_angle

        # angle calcs
        super().__init__(self.get_pos(), *groups, surf=surf, z=z)

    def get_pos(self) -> tuple:
        x = self.center[0] + cos(radians(self.angle)) * self.radius # adjacent
        y = self.center[1] + sin(radians(self.angle)) * self.radius # opponent
        return(x,y)

    def update(self, dt) -> None:
        self.angle += self.speed * dt
        self.rect.center = self.get_pos()
