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
            if hasattr(sprite, "despawn"):
                sprite.kill()
                self.spawn_particle(sprite.rect.center)
                print("player hit by pearl")
            else:
                print("player hit")
