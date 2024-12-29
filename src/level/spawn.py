from enemies.pearl import Pearl
from sprites.particle import Particle

from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from . import Level

def spawn_pearl(self: "Level", pos, direction):
    Pearl(
        pos, self.collision_sprites,
        (self.all_sprites, self.reversible_sprites, self.damage_sprites),
        surf = self.pearl_surface,
        bullet_direction = direction,
        particle_funct = self.spawn_particle
    )

def spawn_particle(self: "Level", pos):
    Particle(pos, self.particle_frames, self.all_sprites)
