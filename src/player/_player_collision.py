from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Player

def collision(self: "Player", axis):
    for sprite in self.collision_sprites:
        if sprite.rect.colliderect(self.rect):
            if axis == "horizontal":
                if (
                    self.rect.right >= sprite.rect.left
                    and int(self.old_rect.right) <= sprite.old_rect.left
                ):
                    self.rect.right = sprite.rect.left
                if (
                    self.rect.left <= sprite.rect.right
                    and int(self.old_rect.left) >= sprite.old_rect.right
                ):
                    self.rect.left = sprite.rect.right
            else:
                if (
                    self.rect.bottom >= sprite.rect.top
                    and int(self.old_rect.bottom) <= sprite.old_rect.top
                ):
                    self.rect.bottom = sprite.rect.top
                if (
                    self.rect.top <= sprite.rect.bottom
                    and int(self.old_rect.top) >= sprite.old_rect.bottom
                ):
                    self.rect.top = sprite.rect.bottom
                self.direction.y = 0

def platform_collision(self: "Player"):
    if not self.timers["platform t/o"]:
        for sprite in self.platform_sprites:
            if sprite.rect.colliderect(self.rect):
                if (
                    self.rect.bottom >= sprite.rect.top
                    and int(self.old_rect.bottom) <= sprite.old_rect.top
                ):
                    self.rect.bottom = sprite.rect.top
                    self.direction.y = 0 if self.direction.y > 0 else self.direction.y

def check_collision_side(self: "Player"):
    collide_rects = [x.rect for x in self.collision_sprites]
    platform_rects = [x.rect for x in self.platform_sprites]
    floor_rect = pygame.Rect(self.rect.bottomleft, (self.rect.width, 2))
    left_rect = pygame.Rect(
        self.rect.topleft + vector(-2, self.rect.height / 4),
        (2, self.rect.h / 2)
    )
    right_rect = pygame.Rect(
        self.rect.topright + vector(0, self.rect.height / 4),
        (2, self.rect.h / 2)
    )

    self.collides_with["floor"] = True if floor_rect.collidelist(collide_rects + platform_rects) >= 0 else False
    self.collides_with["left"] = True if left_rect.collidelist(collide_rects) >= 0 else False
    self.collides_with["right"] = True if right_rect.collidelist(collide_rects) >= 0 else False

    self.platform = None
    for sprite in self.platform_sprites:
        if sprite.rect.colliderect(floor_rect): self.platform = sprite
