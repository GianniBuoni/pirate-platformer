from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Player

def collision(self: "Player", direction):
    for sprite in self.collision_sprites:
        if sprite.rect.colliderect(self.rect):
            if direction == "horizontal":
                if self.rect.right >= sprite.rect.left and self.old_rect.right <= sprite.old_rect.left: # direction.x = 1
                    self.rect.right = sprite.rect.left
                if self.rect.left <= sprite.rect.right and self.old_rect.left >= sprite.old_rect.right: # direction.x = -1
                    self.rect.left = sprite.rect.right
            else:
                if self.rect.bottom >= sprite.rect.top and self.old_rect.bottom <= sprite.old_rect.top: # direction.y = -1
                    self.rect.bottom = sprite.rect.top
                if self.rect.top <= sprite.rect.bottom and self.old_rect.top >= sprite.old_rect.bottom: # direction.y = 1
                    self.rect.top = sprite.rect.bottom
                self.direction.y = 0

def check_collision_side(self: "Player"):
    pass
