from pytmx import TiledMap, TiledObject 
from typing import TYPE_CHECKING

from sprites.items import Item

if TYPE_CHECKING:
    from level import Level

def setup(self: "Level", tmx_map: TiledMap, level_frames):
    for obj in tmx_map.get_layer_by_name("Items"):
        obj: TiledObject = obj

        Item(
            (obj.x + obj.width / 2 , obj.y + obj.height / 2),
            obj.name,
            level_frames["items"][obj.name],
            (self.all_sprites, self.item_sprites)
        )
