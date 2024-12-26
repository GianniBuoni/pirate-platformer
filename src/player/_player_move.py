from typing import TYPE_CHECKING
from ._collision_enum import CollidesWith

if TYPE_CHECKING:
    from . import Player

def move(self: "Player", dt):
    # horizontal movement
    self.hitbox.x += self.direction.x * self.speed * dt
    self.collision("horizontal")

    # vert
    match self.collides_with:
        case CollidesWith.LEFT | CollidesWith.RIGHT:
        # wall slide decreases gravity
            if not self.timers["wjump t/o"]:
                self.direction.y = 0
                self.hitbox.y += self.gravity / 10 * dt
        case _: # normal gravity
            self.direction.y += self.gravity / 2 * dt
            self.hitbox.y += self.direction.y * dt
            self.direction.y += self.gravity / 2 * dt

    self.collision("vertical")

    # platform logic
    if self.platform:
        self.hitbox.center += self.platform.direction * self.platform.speed * dt
    self.platform_collision()

    # jumping logic
    if self.jump:
        match self.collides_with:
            case CollidesWith.FLOOR:
                self.timers["wjump t/o"].activate()
                self.direction.y = -self.jump_distance
            case CollidesWith.LEFT | CollidesWith.RIGHT:
                self.timers["jump t/o"].activate()
                self.direction.y = -self.jump_distance
                self.direction.x = 1 if self.collides_with == CollidesWith.LEFT else -1
        self.jump = False

    self.rect.center = self.hitbox.center
