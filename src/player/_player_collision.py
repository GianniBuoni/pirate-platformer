from player._collision_enum import CollidesWith
from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Player

def collision(self: "Player", axis):
    for sprite in self.collision_sprites:
        if sprite.rect.colliderect(self.hitbox):
            if axis == "horizontal":
                if (
                    self.hitbox.right >= sprite.rect.left
                    and int(self.old_rect.right) <= int(sprite.old_rect.left)
                ):
                    self.hitbox.right = sprite.rect.left
                if (
                    self.hitbox.left <= sprite.rect.right
                    and int(self.old_rect.left) >= int(sprite.old_rect.right)
                ):
                    self.hitbox.left = sprite.rect.right
            else:
                if (
                    self.hitbox.bottom >= sprite.rect.top
                    and int(self.old_rect.bottom) <= int(sprite.old_rect.top)
                ):
                    self.hitbox.bottom = sprite.rect.top
                if (
                    self.hitbox.top <= sprite.rect.bottom
                    and int(self.old_rect.top) >= int(sprite.old_rect.bottom)
                ):
                    self.hitbox.top = sprite.rect.bottom
                self.direction.y = 0

def platform_collision(self: "Player"):
    if not self.timers["platform t/o"]:
        for sprite in self.platform_sprites:
            if sprite.rect.colliderect(self.hitbox):
                if (
                    self.hitbox.bottom >= sprite.rect.top
                    and int(self.old_rect.bottom) <= int(sprite.old_rect.top)
                ):
                    self.hitbox.bottom = sprite.rect.top
                    self.direction.y = 0 if self.direction.y > 0 else self.direction.y

def check_collision_side(self: "Player"):
    collide_rects = [x.rect for x in self.collision_sprites]
    platform_rects = [x.rect for x in self.platform_sprites]
    floor_rect = pygame.Rect(self.hitbox.bottomleft, (self.hitbox.width, 2))
    left_rect = pygame.Rect(
        self.hitbox.topleft + vector(-2, self.hitbox.height / 4),
        (2, self.hitbox.h / 2)
    )
    right_rect = pygame.Rect(
        self.hitbox.topright + vector(0, self.hitbox.height / 4),
        (2, self.hitbox.h / 2)
    )

    if floor_rect.collidelist(collide_rects + platform_rects) >= 0:
        self.collides_with = CollidesWith.FLOOR
    elif left_rect.collidelist(collide_rects) >= 0:
        self.collides_with = CollidesWith.LEFT
    elif right_rect.collidelist(collide_rects) >= 0:
        self.collides_with = CollidesWith.RIGHT
    else: self.collides_with = CollidesWith.AIR

    self.platform = None
    for sprite in self.platform_sprites:
        if (
            sprite.rect.colliderect(floor_rect)
            and hasattr(sprite, "moving")
        ):
            self.platform = sprite
