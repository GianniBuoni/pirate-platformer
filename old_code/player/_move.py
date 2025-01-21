from typing import TYPE_CHECKING
from ._enums import CollidesWith

if TYPE_CHECKING:
    from . import Player

def move(self: "Player", dt):
    # horizontal movement
    self.hitbox.x += self.direction.x * self.speed * dt
    self.collision("horizontal")

    # vert
    match self.check_collision_side():
        case CollidesWith.FLOOR | CollidesWith.AIR:
            self.reset_gravity(dt)
        case CollidesWith.LEFT | CollidesWith.RIGHT:
            if not self.timers["wjump t/o"]:
                self.direction.y = 0
                self.hitbox.y += self.gravity / 10 * dt
            else:
                self.reset_gravity(dt)

    self.collision("vertical")

    # platform logic
    if self.platform():
        self.hitbox.center += self.platform().direction * self.platform().speed * dt # pyright: ignore
    self.platform_collision()

    # jumping logic
    if self.jump:
        match self.check_collision_side():
            case CollidesWith.FLOOR:
                self.timers["wjump t/o"].activate()
                self.direction.y = -self.jump_distance
            case CollidesWith.LEFT | CollidesWith.RIGHT:
                if not self.timers["wjump t/o"]:
                    self.timers["jump t/o"].activate()
                    self.direction.x = 1 if self.check_collision_side() == CollidesWith.LEFT else -1
                self.direction.y = -self.jump_distance
        self.jump = False

    self.rect.center = self.hitbox.center

def reset_gravity(self: "Player", dt):
    self.direction.y += self.gravity / 2 * dt
    self.hitbox.y += self.direction.y * dt
    self.direction.y += self.gravity / 2 * dt
