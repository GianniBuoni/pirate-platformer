__all__ = ["Level"]

from settings import *
from lib.groups import *
from level.setup_functions import *
from player import Player

class Level:
    def __init__(self, tmx_map, level_frames, data):
        self.display_surface = pygame.display.get_surface()
        self.data = data

        # level_info:
        self.level_w = tmx_map.width * TILE_SIZE
        self.level_h = tmx_map.height * TILE_SIZE

        # groups
        self.all_sprites = AllSprites()
        self.collision_sprites = CollisionSprites()
        self.platform_sprites = pygame.sprite.Group()
        self.damage_sprites = pygame.sprite.Group()
        self.item_sprites = pygame.sprite.Group()
        self.reversible_sprites = pygame.sprite.Group()

        # sprites
        self.pearl_surface = level_frames["pearl"]
        self.particle_frames = level_frames["particle"]
        self.player = Player(
            (0, 0), level_frames["player"],
            self.collision_sprites,
            self.platform_sprites,
            self.all_sprites
        )

        self.setup(tmx_map, level_frames)

    def setup(self, tmx_map, level_frames):
        args = (self, tmx_map, level_frames)

        tiles.setup_tiles(*args)
        bg.setup(*args)
        objects.setup(*args)
        moving_objects.setup(*args)
        enemies.setup(*args)
        items.setup(*args)
        water.setup(*args)

    from .spawn import spawn_pearl, spawn_particle
    from .collisions import check_collisions, get_item
    from .constraints import check_constraints

    def run(self, dt):
        self.display_surface.fill("black")
        self.all_sprites.update(dt)
        self.check_collisions()
        self.all_sprites.draw(self.player.hitbox.center)
        self.check_constraints()
