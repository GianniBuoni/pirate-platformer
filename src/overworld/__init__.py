__all__ = ["Overworld"]

from lib.groups import AllSprites
from settings import *
from overworld.setup_functions import *
from lib.data import Data

class Overworld():
    def __init__(self, tmx_map, overworld_frames, data: Data) -> None:
        self.display_surface = pygame.display.get_surface()
        self.data = data

        # overworld data
        self.cam_offset = vector(-500, -300)

        # groups
        self.all_sprites = AllSprites()

        # events
        self.setup(tmx_map, overworld_frames)

    def setup(self, tmx_map, overworld_frames):
        args = (self, tmx_map, overworld_frames)

        tiles.setup(*args)
        water.setup(*args)
        objects.setup(*args)
        nodes.setup(*args)

    def run(self, dt):
        self.all_sprites.update(dt)
        self.all_sprites.draw_overworld(self.cam_offset)
