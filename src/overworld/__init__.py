__all__ = ["Overworld"]

from typing import Union
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
        self.paths = {}
        self.current_path = []
        self.start_point: Union[tuple[float, float], None] = None

        # sprites
        self.all_sprites = AllSprites()
        self.node_sprites = pygame.sprite.Group()
        self.icon = Icon(self.all_sprites, frames=overworld_frames["icon"])

        # events
        self.setup(tmx_map, overworld_frames)

    def setup(self, tmx_map, overworld_frames):
        self.get_paths(tmx_map)
        args = (self, tmx_map, overworld_frames)

        tiles.setup(*args)
        water.setup(*args)
        objects.setup(*args)
        nodes.setup(*args)

    def input(self):
        keys = pygame.key.get_pressed()
        valid_inputs = self.availabe_inputs()[0]

        if keys[pygame.K_LEFT] and "left" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["left"][1:]
        if keys[pygame.K_RIGHT] and "right" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["right"][1:]
        if keys[pygame.K_UP] and "up" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["up"][1:]
        if keys[pygame.K_DOWN] and "down" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["down"][1:]

    from .movement import move_icon, offset_camera, pivot_path_points
    from .pathing import get_paths, availabe_inputs, availabe_paths

    def run(self, dt):
        self.display_surface.fill("black")
        self.input()
        self.move_icon()
        self.offset_camera(self.icon.rect.center)
        self.all_sprites.update(dt)
        self.all_sprites.draw_overworld(self.cam_offset)
