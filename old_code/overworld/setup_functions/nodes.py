from settings import *
from sprites import *

from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from overworld import Overworld
    import pytmx

def setup(self: "Overworld", tmx_map: "pytmx.TiledMap", frames):
    for obj in tmx_map.get_layer_by_name("Nodes"):
        obj: "pytmx.TiledObject" = obj

        # render unlocked levels
        if obj.properties["stage"] <= self.data.unlocked_levels:
            node.Node(
                (obj.x, obj.y), self.all_sprites, self.node_sprites,
                surf= frames["path"]["node"],
                tmx_props=obj.properties
            )

        # render player on current level
        match obj.properties["stage"]:
            case self.data.current_level:
                self.icon.rect.topleft = (obj.x, obj.y)
