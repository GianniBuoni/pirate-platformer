from settings import *
from sprites.sprites import Sprite

from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from overworld import Overworld
    import pytmx

def setup(self: "Overworld", tmx_map: "pytmx.TiledMap", _):
    for layer in ["main", "top"]:
        for x, y, surf in tmx_map.get_layer_by_name(layer).tiles():
            Sprite(
                (x * TILE_SIZE, y * TILE_SIZE),
                self.all_sprites,
                surf= surf,
                z= Z_LAYERS["bg tiles"]
            )
