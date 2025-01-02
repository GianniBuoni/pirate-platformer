from settings import *
from .sprites import Sprite

class Node(Sprite):
    def __init__(self, pos, *groups, surf) -> None:
        super().__init__(pos, *groups, surf=surf, z= Z_LAYERS["path"])
        self.rect.center = (
            pos[0] + (TILE_SIZE / 2),
            pos[1] + (TILE_SIZE / 2)
        )
