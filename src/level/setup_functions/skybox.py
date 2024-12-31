import pytmx
from settings import *
from sprites.sprites import Sprite

from typing import TYPE_CHECKING

if TYPE_CHECKING: from level import Level

def setup(self: "Level", tmx_map: pytmx.TiledMap, level_frames):
    bg_tile = self.level_data["bg"]
    if bg_tile:

        top_limit = self.level_data["top_limit"]
        extra_rows = int(top_limit / TILE_SIZE) + 1 if top_limit else 0

        cols, rows, = tmx_map.width, tmx_map.height + extra_rows
        for col in range(cols):
            for row in range(rows):
                x = col * TILE_SIZE
                y = row * TILE_SIZE - extra_rows * TILE_SIZE
                Sprite((x, y), self.all_sprites, surf=level_frames["bg"][bg_tile], z=-1)
    else:
        self.sky = True
