from settings import *
from sprites import *
from sprites.icon import Icon

from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from overworld import Overworld
    import pytmx

def setup(self: "Overworld", tmx_map: "pytmx.TiledMap", frames):
    for obj in tmx_map.get_layer_by_name("Nodes"):
        obj: "pytmx.TiledObject" = obj
        print(f"{obj.name} = {obj.properties}")

        node.Node((obj.x, obj.y), self.all_sprites, surf= frames["path"]["node"])

        # render player on current level
        match obj.properties["stage"]:
            case self.data.current_level:
                Icon(self.all_sprites, pos=(obj.x, obj.y), frames=frames["icon"])
