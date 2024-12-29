from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Level

def check_collisions(self: "Level"):
    # item and player
    if self.item_sprites.sprites():
        for sprite in self.item_sprites:
            collision = sprite.rect.colliderect(self.player.hitbox)
            if collision:
                sprite.kill()
                self.spawn_particle(sprite.rect.center)
                print(f"item collected: {sprite.item_type}")

    # pearl and player
    for sprite in self.damage_sprites:
        collision = sprite.rect.colliderect(self.player.hitbox)
        if collision:
            self.player.damaged()
            if hasattr(sprite, "particle"):
                sprite.kill()
                self.spawn_particle(sprite.rect.center)

    # attack collisions
    for target in self.reversible_sprites:
        collision = target.rect.colliderect(self.player.rect)
        facing_target = (
            self.player.rect.center < target.rect.center and self.player.facing_right
            or self.player.rect.center > target.rect.center and not self.player.facing_right
        )
        if collision and self.player.attacking and facing_target:
            target.reverse()
