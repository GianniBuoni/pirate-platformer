from settings import *
from sprites.animated import AnimatedSprite

from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from overworld import Overworld
    import pytmx

def setup(self: "Overworld", tmx_map: "pytmx.TiledMap", frames):
    for col in range(tmx_map.width):
        for row in range(tmx_map.height):
            x = col * TILE_SIZE
            y = row * TILE_SIZE

            AnimatedSprite(
                (x, y),
                frames["water"],
                self.all_sprites,
                z= Z_LAYERS["bg"]
            )
