from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Player

def move(self: "Player", dt):
    # horizontal movement
    self.hitbox.x += self.direction.x * self.speed * dt
    self.collision("horizontal")

    # vert
    any_wall_collide = (
        any((self.collides_with["left"], self.collides_with["right"]))
        and not self.timers["wjump t/o"]
    )

    # wall slide decreases gravity
    if not self.collides_with["floor"] and any_wall_collide:
        self.direction.y = 0
        self.hitbox.y += self.gravity / 10 * dt
    else: # normal gravity
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
        if self.collides_with["floor"]:
            self.timers["wjump t/o"].activate()
            self.direction.y = -self.jump_distance
        elif any_wall_collide:
            self.timers["jump t/o"].activate()
            self.direction.y = -self.jump_distance
            self.direction.x = 1 if self.collides_with["left"] else -1
        self.jump = False

    self.rect.center = self.hitbox.center
