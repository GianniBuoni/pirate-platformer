__all__ = ["Overworld"]

from settings import *
from overworld.setup_functions import *

class Overworld():
    def __init__(self, tmx_map, data, overworld_frames) -> None:
        self.display_surface = pygame.display.get_surface()
        self.data = data

        # groups
        self.all_sprites = pygame.sprite.Group()

        # events
        self.setup(tmx_map, overworld_frames)

    def setup(self, tmx_map, overworld_frames):
        args = (self, tmx_map, overworld_frames)

        tiles.setup(*args)

    def run(self, dt):
        self.all_sprites.update()
        self.all_sprites.draw(self.display_surface)
