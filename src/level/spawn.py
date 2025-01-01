from random import randint, choice

from settings import *
from enemies.pearl import Pearl
from sprites import *

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
    particle.Particle(pos, self.particle_frames, self.all_sprites)

def spawn_cloud(self: "Level"):
    surface = choice(self.clouds["small"])

    x = randint(
        int(self.level_w + surface.width / 2),
        int(self.level_w + WINDOW_WIDTH / 2)
    )
    y = randint(0 , int(self.level_data["horizon_line"] / 2))

    cloud.Cloud((x, y), self.all_sprites, surf = surface)
