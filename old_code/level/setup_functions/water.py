import pytmx
from settings import *
from sprites import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from level import Level

def setup(self: "Level", tmx_map: pytmx.TiledMap, level_frames):
    for obj in tmx_map.get_layer_by_name("Water"):
        obj: pytmx.TiledObject = obj
        rows = int(obj.height / TILE_SIZE)
        cols = int(obj.width / TILE_SIZE)

        # TO DO: look into a way to avoid nested for loop later
        for row in range(rows):
            for col in range(cols):
                x = obj.x + col * TILE_SIZE
                y = obj.y + row * TILE_SIZE
                
                if row == 0:
                    animated.AnimatedSprite(
                        (x, y), level_frames["water_top"], self.all_sprites, z=Z_LAYERS["fg"]
                    )
                else:
                    sprites.Sprite(
                        (x, y), self.all_sprites, surf=level_frames["water_body"], z=Z_LAYERS["fg"]
                    )
