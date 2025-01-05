__all__ = ["Overworld"]

from typing import Union
from settings import *
from overworld.setup_functions import *

from lib.groups import AllSprites
from lib.data import Data
from sprites.icon import Icon

class Overworld():
    def __init__(self, tmx_map, overworld_frames, data: Data, switch_stage) -> None:
        self.display_surface = pygame.display.get_surface()
        self.data = data
        self.switch_stage = switch_stage

        # overworld data
        self.cam_offset = vector()
        self.paths = {}
        self.current_path = []
        self.start_point: Union[tuple[float, float], None] = None
        self.can_input = True

        # sprites
        self.all_sprites = AllSprites()
        self.node_sprites = pygame.sprite.Group()
        self.icon = Icon(self.all_sprites, frames=overworld_frames["icon"])

        # events
        self.setup(tmx_map, overworld_frames)
        self.create_path(overworld_frames)

    def setup(self, tmx_map, overworld_frames):
        self.get_paths(tmx_map)
        args = (self, tmx_map, overworld_frames)

        tiles.setup(*args)
        water.setup(*args)
        objects.setup(*args)
        nodes.setup(*args)

    def check_current_node(self):
        node = pygame.sprite.spritecollide(self.icon, self.node_sprites, False)
        if node:
            self.data.current_level = node[0].id
            self.can_input = True

    from ._input import input
    from .movement import move_icon, offset_camera, pivot_path_points
    from .pathing import get_paths, availabe_inputs, availabe_paths, create_path

    def run(self, dt):
        self.display_surface.fill("black")
        self.input()
        self.move_icon()
        self.offset_camera(self.icon.rect.center)
        self.all_sprites.update(dt)
        self.all_sprites.draw_overworld(self.cam_offset)
