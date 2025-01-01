__all__ = ["Level"]

from settings import *
from lib.groups import *
from lib.timer import Timer
from level.setup_functions import *
from player import Player

class Level:
    def __init__(self, tmx_map, level_frames, data):
        # game data
        self.display_surface = pygame.display.get_surface()
        self.data = data

        # level_data:
        self.level_w = tmx_map.width * TILE_SIZE
        self.level_h = tmx_map.height * TILE_SIZE
        self.level_data = tmx_map.get_layer_by_name("Data")[0].properties
        self.cam_offset = vector()
        self.sky, self.sky_offset = False, 0

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
        self.clouds = {
            "small": level_frames["cloud_small"],
            "large": level_frames["cloud_large"]
        }
        self.player = Player(
            (0, 0), level_frames["player"],
            self.collision_sprites,
            self.platform_sprites,
            self.all_sprites
        )

        # events
        self.setup(tmx_map, level_frames)
        if self.sky:
            self.cloud_timer = Timer(3000, func= self.spawn_cloud, repeat= True)
            self.cloud_timer.activate()

    def setup(self, tmx_map, level_frames):
        args = (self, tmx_map, level_frames)

        skybox.setup(*args)
        tiles.setup_tiles(*args)
        bg.setup(*args)
        objects.setup(*args)
        moving_objects.setup(*args)
        enemies.setup(*args)
        items.setup(*args)
        water.setup(*args)

    from ._skybox import draw_skybox
    from .camera import offset_camera
    from .collisions import check_collisions, get_item
    from .constraints import check_constraints
    from .spawn import spawn_pearl, spawn_particle, spawn_cloud

    def run(self, dt):
        if self.sky:
            self.cloud_timer.update()
            self.draw_skybox(dt)
        else:
            self.display_surface.fill("black")
        self.check_constraints()
        self.offset_camera(self.player.hitbox.center)
        self.all_sprites.update(dt)
        self.check_collisions()
        self.all_sprites.draw(self.cam_offset)
