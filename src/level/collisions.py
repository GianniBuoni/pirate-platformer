from settings import *
from sprites.particle import Particle
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Level

def item_player(self: "Level"):
    if self.item_sprites.sprites() and self.player:
        for sprite in self.item_sprites:
            collision = sprite.rect.colliderect(self.player.hitbox)
            if collision:
                sprite.kill()
                Particle(sprite.rect.center, self.particle_frames, self.all_sprites)
                print(f"item collected: {sprite.item_type}")
