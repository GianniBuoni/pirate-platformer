__all__ = ["Level"]

from settings import *
from lib.groups import *
from level.setup_functions import *

class Level:
    def __init__(self, tmx_map, level_frames):
        self.display_surface = pygame.display.get_surface()
        self.player = None

        # groups
        self.all_sprites = AllSprites()
        self.collision_sprites = CollisionSprites()
        self.platform_sprites = pygame.sprite.Group()
        self.damage_sprites = pygame.sprite.Group()

        self.setup(tmx_map, level_frames)

    def setup(self, tmx_map, level_frames):
        args = (self, tmx_map, level_frames)

        tiles.setup_tiles(*args)
        bg.setup(*args)
        objects.setup(*args)
        moving_objects.setup(*args)
        enemies.setup(*args)
        items.setup(*args)

    def run(self, dt):
        self.display_surface.fill("black")
        self.all_sprites.update(dt)
        self.all_sprites.draw(self.player.hitbox.center) # pyright: ignore # self.player defined in objects.setup()

