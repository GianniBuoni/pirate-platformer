__all__ = ["Overworld"]

from settings import *
from overworld.setup_functions import *

from lib.groups import AllSprites
from lib.data import Data
from sprites.icon import Icon

class Overworld():
    def __init__(self, tmx_map, overworld_frames, data: Data) -> None:
        self.display_surface = pygame.display.get_surface()
        self.data = data

        # overworld data
        self.cam_offset = vector()

        # sprites
        self.all_sprites = AllSprites()
        self.icon = Icon(self.all_sprites, frames=overworld_frames["icon"])

        # events
        self.setup(tmx_map, overworld_frames)

    def setup(self, tmx_map, overworld_frames):
        args = (self, tmx_map, overworld_frames)

        tiles.setup(*args)
        water.setup(*args)
        objects.setup(*args)
        nodes.setup(*args)

    from .camera import offset_camera

    def run(self, dt):
        self.offset_camera(self.icon.rect.center)
        self.all_sprites.update(dt)
        self.all_sprites.draw_overworld(self.cam_offset)
